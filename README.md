# go-cloudconvert
File Convert Tool using CloudConvert API V2

## Required
* go >= 1.16

## Supported Format
 +200 Formats check for https://cloudconvert.com
 * Archive
 > 7z RAR ZIP ...
 
 * Audio         
 > AAC FLAC M4A MP3 WAV ...
 * Cad           
 > DWG DXF
 * Document      
 > DOC HTML TXT PDF ...
 * EBook        
 > EPUB PDB ...
 * Font         
 > TTF OTF ...
 * Image         
 > HEIC GIF JPEG PNG TIFF ...
 * Presentation  
 > KEY PPT ...
 * Spreadsheet   
 > CSV NUMBERS XLS XLSX ...
 * Vector        
 > AI SVG ...
 * Video         
 > MOV FLV MP4 WEBM ...
 
 

## Usage
* Get API Token from https://cloudconvert.com
* Make go-cloudconvert/convert/config.ini  and Set API token
### option 
* -i : input filepath like ./sample.jpg
* -f : output format like pdf
* -l : print convertible formats corresponding to input
* -h : print all supported formats
