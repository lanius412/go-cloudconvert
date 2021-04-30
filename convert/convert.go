package convert

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

//go:embed config.ini
var token string

type Tasks struct {
	Bearer        string
	ImportTaskID  string
	ConvertTaskID string
	ExportTaskID  string
}

var tasks Tasks

//Do do convert process
func Do(filePath, outputFormat string) (string, error) {
	var err error
	bearer := "Bearer " + token
	tasks.Bearer = bearer

	encodedFile, fileName := encodeBase64(filePath)

	err = tasks.uploadFile(encodedFile, fileName)
	if err != nil {
		log.Fatalln("import task error: ", err)
	}

	err = tasks.convertFile(outputFormat)
	if err != nil {
		log.Fatalln("convert task error: ", err)
	}

	err = tasks.exportFile()
	if err != nil {
		log.Fatalln("export task error: ", err)
	}

	downloadUrl, pdfFileName := tasks.doneTask()
	if downloadUrl == "" || pdfFileName == "" {
		return "", errors.New("download url can't find")
	}

	err = tasks.deleteTasks()
	if err != nil {
		log.Fatalln("delete tasks error: ", err)
	}

	err = downlaodFile(downloadUrl, pdfFileName)
	if err != nil {
		log.Fatalln("download file error: ", err)
	}

	return pdfFileName, err
}

//downloadFile download a pdf file from cloudconvert storage url
func downlaodFile(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

//deleteTasks delete import, convert and export task
func (tasks *Tasks) deleteTasks() error {
	taskIDs := []string{tasks.ImportTaskID, tasks.ConvertTaskID, tasks.ExportTaskID}
	baseUrl, err := url.Parse("https://api.cloudconvert.com/v2/tasks/")
	if err != nil {
		return err
	}
	for _, id := range taskIDs {
		idUrl, err := url.Parse(id)
		if err != nil {
			return err
		}
		endpoint := baseUrl.ResolveReference(idUrl).String()
		req, err := http.NewRequest("DELETE", endpoint, nil)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", tasks.Bearer)
		client := &http.Client{}
		_, err = client.Do(req)
		if err != nil {
			return err
		}
	}
	return err
}

//doneTask return request url whose body has a converted file and converted file name
func (tasks *Tasks) doneTask() (string, string) {
	var downloadUrl, fileName string
	for {
		body, err := tasks.checkTasks()
		if err != nil {
			log.Fatalln("check task error: ", err)
		}
		defer body.Close()
		listResp := new(ListTaskResp)
		if err := json.NewDecoder(body).Decode(&listResp); err != nil {
			log.Fatalln("read check tasks response error: ", err)
		}
		if len(listResp.Data[0].Result.Files) != 0 {
			fileName = listResp.Data[0].Result.Files[1].FileName
			downloadUrl = listResp.Data[0].Result.Files[1].Url
			break
		}
	}
	fmt.Println("Convert Finish")
	return downloadUrl, fileName
}

//checkTasks return tasks response
func (tasks *Tasks) checkTasks() (io.ReadCloser, error) {
	endpoint := "https://api.cloudconvert.com/v2/tasks"
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", tasks.Bearer)
	req.URL.Query().Set("filter[operation]", "export/url")
	req.URL.Query().Set("page", "true")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, err
}

//postRequest send  each POST REQUEST
func postRequest(endpoint string, reqJson []byte) (io.ReadCloser, error) {
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqJson))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", tasks.Bearer)
	req.Header.Add("Content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, err
}

//exportFile create export/url task
func (tasks *Tasks) exportFile() error {
	endpoint := "https://api.cloudconvert.com/v2/export/url"
	exportReq := new(ExportTaskReq)
	exportReq.Input[0] = tasks.ImportTaskID
	exportReq.Input[1] = tasks.ConvertTaskID
	reqJson, err := json.Marshal(exportReq)
	if err != nil {
		return err
	}
	body, err := postRequest(endpoint, reqJson)
	if err != nil {
		log.Fatalln("post request error: ", err)
	}
	defer body.Close()
	exportResp := new(ExportTaskResp)
	if err := json.NewDecoder(body).Decode(&exportResp); err != nil {
		return err
	}
	fmt.Println("Operation:", exportResp.Data.Operation)
	tasks.ExportTaskID = exportResp.Data.ID
	return err
}

//convertFile create convert Task and convert uploaded file
func (tasks *Tasks) convertFile(outputFormat string) error {
	endpoint := "https://api.cloudconvert.com/v2/convert"
	convertReq := new(ConvertTaskReq)
	convertReq.Input = tasks.ImportTaskID
	//convertReq.OutputFormat = "pdf"
	convertReq.OutputFormat = outputFormat
	reqJson, err := json.Marshal(convertReq)
	if err != nil {
		return err
	}
	body, err := postRequest(endpoint, reqJson)
	if err != nil {
		log.Fatalln("post request error: ", err)
	}
	defer body.Close()
	convertResp := new(ConvertTaskResp)
	if err := json.NewDecoder(body).Decode(&convertResp); err != nil {
		return err
	}
	fmt.Println("Operation:", convertResp.Data.Operation)
	tasks.ConvertTaskID = convertResp.Data.ID
	return err
}

//uploadFile create import/base64 Task and upload encoded file
func (tasks *Tasks) uploadFile(encodedFile, fileName string) error {
	endpoint := "https://api.cloudconvert.com/v2/import/base64"
	importReq := new(ImportTaskReq)
	importReq.File = encodedFile
	importReq.FileName = fileName
	reqJson, err := json.Marshal(importReq)
	if err != nil {
		return err
	}
	body, err := postRequest(endpoint, reqJson)
	if err != nil {
		log.Fatalln("post request error: ", err)
	}
	defer body.Close()
	importResp := new(ImportTaskResp)
	if err := json.NewDecoder(body).Decode(&importResp); err != nil {
		return err
	}
	fmt.Println("Operation:", importResp.Data.Operation)
	tasks.ImportTaskID = importResp.Data.ID
	return err
}

//encodeBase64 return a byte string file which encoded by Base64 and file name
func encodeBase64(filePath string) (string, string) {
	fileByte, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
	}
	encodedFile := base64.StdEncoding.EncodeToString(fileByte)
	return encodedFile, filepath.Base(filePath)
}
