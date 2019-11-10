package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)
const (
	filepath = "/opt/lampp/htdocs/Dashboard/assets/data/"
)

func validateApi(w http.ResponseWriter, r *http.Request) bool{
	r.ParseForm()                     // Parses the request body
	api := r.Form.Get("api_key") // x will be "" if parameter is not set
	if api == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falta api-key"))
		return false
	}
	if _, ok := apis[api]; !ok {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Api-key invalida"))
		return false
	}
	return true
}
func uploadDataHandler(w http.ResponseWriter, r *http.Request){
	if !validateApi(w,r){
		return
	}
	var Buf bytes.Buffer
	file, _, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	id := apis[r.Form.Get("api_key")] // Obtenemos el ID de al antena que reporta
	ts, _ := strconv.ParseInt(r.Form.Get("ts"), 10, 64)
	timestamp := time.Unix(ts, 0)
	filename := timestamp.Format("02-01-2006-15-04-05") // Formato para el nombre del archivo
	filename = filename+"_"+strconv.Itoa(id)+".json" // Nombre final
	fmt.Println(filename)

	io.Copy(&Buf, file)
	f, err := os.Create(filepath+filename)
	if err != nil {
		panic(err.Error())
		fmt.Println("Error al crear el archivo",filename)
		return
	}
	_,err = f.Write(Buf.Bytes())
	if err !=nil{
		panic(err.Error())
		fmt.Println("Error al escribir en el archivo")
	}
	defer f.Close()

	newReport(id, filepath+filename,timestamp)

}
func notificationHandler(w http.ResponseWriter, r *http.Request)  {
	if !validateApi(w,r){
		return
	}
	fmt.Println("NOTIFICAR :3")
	//sendNotification()
}
