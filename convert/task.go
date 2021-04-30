package convert

type ImportTaskReq struct {
	File     string `json:"file"`
	FileName string `json:"filename"`
}
type ImportTaskResp struct {
	Data struct {
		ID        string `json:"id"`
		Operation string `json:"operation"`
		Status    string `json:"status"`
	} `json:"data"`
}

type ConvertTaskReq struct {
	Input        string `json:"input"`
	InputFormat  string `json:"input_format"`
	OutputFormat string `json:"output_format"`
}
type ConvertTaskResp struct {
	Data struct {
		ID        string `json:"id"`
		Operation string `json:"operation"`
		Status    string `json:"status"`
	} `json:"data"`
}

type ExportTaskReq struct {
	Input [2]string `json:"input"`
}
type ExportTaskResp struct {
	Data struct {
		ID        string `json:"id"`
		Operation string `json:"operation"`
		Status    string `json:"status"`
	} `json:"data"`
}

type ListTaskResp struct {
	Data []struct {
		Result struct {
			Files []struct {
				FileName string `json:"filename"`
				Url      string `json:"url"`
			} `json:"files"`
		} `json:"result"`
	} `json:"data"`
}
