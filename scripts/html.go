package necrotheif

import (
	"net/http"
	"html/template"
	"io/ioutil"
	"bytes"
	"os"
	"encoding/json"
	"errors"
)

var htmlTemplate *template.Template
var pages map[string]BasicHtml
var styleSheet template.CSS
var javascript template.JS

func loadFiles() error{

	loadColors(backgroundColor, foregroundColor, linkColor)

	cssTemplate, err := template.ParseFiles(templatesFolder+"/"+basicTemplateName+".css")
	if(err!=nil){return err}
	var cssBuff bytes.Buffer
	err = cssTemplate.Execute(&cssBuff, cssBasic)
	if(err!=nil){return errors.New("Error Reading CSS: <br/>"+err.Error())}
	styleSheet = template.CSS(cssBuff.String())

	jsTemplates, err := loadTemplates(jsFolder)
	if(err!=nil){return err}
	var jsBuff bytes.Buffer
	err = jsTemplates.Execute(&jsBuff, nil)
	if(err!=nil){return err}
	javascript = template.JS(jsBuff.String())

	htmlTemplate, err = template.ParseFiles(templatesFolder+"/"+basicTemplateName+".html")
	if(err!=nil){return err}

	pages, err = loadFolder(jsonFolder)
	if(err!=nil){return err}
	subPages, err := loadSubFolders(jsonFolder)
	consoleLog(err)
	if(err!=nil){return err}
	for folder,temp := range subPages{pages[folder]=temp}

	return nil
	
}

func loadSubFolders(dir string) (map[string]BasicHtml, error){
	files, err := ioutil.ReadDir(dir)
	if(err!=nil) {return nil, errors.New("Error Reading Directory: "+dir+"<br/>"+err.Error())}
	curPages := make(map[string]BasicHtml)
	for _,file := range files{
		if(file.IsDir()){
			newPages, err := loadFolder(dir+"/"+file.Name())
			for folder,page := range newPages{curPages[folder]=page}
			if(err!=nil) {return nil, err}
			newPages, err = loadSubFolders(dir+"/"+file.Name())
			if(err!=nil) {return nil, err}
			for folder,page := range newPages{curPages[folder]=page}
		}
	}
	return curPages, nil
}

func loadFolder(dir string) (map[string]BasicHtml, error){
	files, err := ioutil.ReadDir(dir)
	if(err!=nil) {return nil, errors.New("Error Reading Directory: "+dir+"<br/>"+err.Error())}
	pages := make(map[string]BasicHtml)
	for _, fileInfo := range files{
		if(!fileInfo.IsDir()){
			file, err := os.Open(dir+"/"+fileInfo.Name())
			if(err!=nil) {return nil, errors.New("Error Opening file: "+dir+"/"+fileInfo.Name()+"<br/>"+err.Error())}
			data := make([]byte, fileInfo.Size())
			numRead, err := file.Read(data)
			if(err!=nil && numRead==0) {return nil, errors.New("Error Reading file: "+dir+"/"+fileInfo.Name()+"<br/>"+err.Error())}
			var htmlPage BasicHtml
			err = json.Unmarshal(data, &htmlPage)
			if(err!=nil) {return nil, errors.New("Error Reading json: "+dir+"/"+fileInfo.Name()+"<br/>"+err.Error())}
			htmlPage.WebTitle = websiteTitle
			htmlPage.Author = author
			htmlPage.StyleSheet = styleSheet
			htmlPage.Javascript = javascript
			pages[file.Name()] = htmlPage
		}
	}
	return pages, nil
}

func loadTemplates(dir string) (*template.Template, error){
	files, err := ioutil.ReadDir(dir)
	if(err!=nil) {return nil, errors.New("Error Reading Directory: "+dir+"<br/>"+err.Error())}
	names := make([]string, 0)
	for _, file := range files{
		if(!file.IsDir()){
			names = append(names, dir+"/"+file.Name())
		}
	}
	return template.ParseFiles(names...)
}

func loadPage(w http.ResponseWriter, path string, page string){
	
	if(page==""){page="about"}
	if _, exists := pages[jsonFolder+path+"/"+page+".json"]; !exists {
		loadError(w, errors.New("Page Not Found!"))
	} else {
		err := htmlTemplate.Execute(w, pages[jsonFolder+path+"/"+page+".json"])
		if err != nil {loadError(w, err)}
	}
}

func loadError(w http.ResponseWriter, err error){
	errPage := pages[jsonFolder+"/"+errorPage+".json"]
	errPage.Message = template.HTML("Error:<br/>"+err.Error())
	htmlTemplate.Execute(w, errPage)
}