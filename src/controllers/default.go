package controllers

import (
	"html/template"
	"log"
	//"models"
	//"strings"
	"net/http"
	"fmt"
    "os"
	"io"
	"io/ioutil"
	"os/exec"
)

const(
	upload_dir = "./src/tmp/input/" //don't forget the last slash
)


////////////////////////////////////////////////////////////////////////////////////////////////
//just want to make it looks easier, ignore all error checking temporarily


func ProcessRequest(rw http.ResponseWriter, req *http.Request){
	outputFileName := ProcessedFile(UploadedFile(rw,req))
	//fmt.Println(outputFileName)
	outputContent, _ := ioutil.ReadFile(outputFileName)
	fmt.Fprintln(rw,string(outputContent))
}

func UploadedFile(w http.ResponseWriter, r *http.Request)string{
	inputFile, header, _ := r.FormFile("upload-file")
	defer inputFile.Close()
	uploadedFile, _ := os.OpenFile(upload_dir + header.Filename, os.O_CREATE|os.O_WRONLY, 0660)
	defer uploadedFile.Close()
	io.Copy(uploadedFile, inputFile)
	uploadedFileName := uploadedFile.Name()
	return uploadedFileName
}

func ProcessedFile(uploadedFile string)string{
	executablePath := "./src/executable/test.exe"
	argv := []string{uploadedFile}
	cmd := exec.Command(executablePath, argv...)
	output, _ := cmd.Output()
	outputFileName := string(output)
	//fmt.Println(string(output))
	return outputFileName
}
////////////////////////////////////////////////////////////////////////////////////////////////

func HomeController(rw http.ResponseWriter, req *http.Request) {

	renderHomepage(rw, req)
}

func renderHomepage(rw http.ResponseWriter, req *http.Request) {
	
	// grab the homepage from views 
	homepage, err := template.ParseFiles("src/views/html/homepage.html")
	
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// send it to the browser
	homepage.Execute(rw, homepage)
}

/*
func Upload(rw http.ResponseWriter, req *http.Request) {
  
	// "upload-file" is from the POST method of the form on the web page
	uploadfile, header, err := req.FormFile("upload-file")

	if err != nil {
		fmt.Fprintln(rw, err)
		return
	}

	// fancy clean way to close file after the function returns
	defer uploadfile.Close()
	
	// creates a file for the newly uplaoded file to be copied into, from the POST

	//serverFile, err := os.Create("/tmp/uploadedfile") //"yourfilename.ext"
	
	//serverFile, err := os.Create(upload_dir + header.Filename) //fine

	serverFile, err := os.OpenFile(upload_dir + header.Filename, os.O_CREATE|os.O_WRONLY, 0660) //fine
	
	if err != nil {
		fmt.Fprintln(rw, err)
		return
	}
	
	defer serverFile.Close()
	
	// writes to the serverFile from the POST 
	_, err = io.Copy(serverFile, uploadfile)
	
	if err != nil {
		fmt.Fprintln(rw, err)
	}

	//fmt.Fprintf(rw, "File uploaded successfully : ")
    //fmt.Fprintf(rw, header.Filename)	
}
*/