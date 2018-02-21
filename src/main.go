package main

import (
	"./contenttypes"
	"fmt"
	"github.com/golang/gddo/httputil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const buffersize = 1024 * 1024

func handler(rw http.ResponseWriter, r *http.Request) {
	strDelay := r.URL.Query().Get("delay")
	strSize := r.URL.Query().Get("size")
	strRaw := r.URL.Query().Get("raw")

	sizeInByte := 42
	if len(strSize) > 0 {
		faktor := 1
		if len(strSize) > 1 {
			if strSize[len(strSize)-1] == 'k' { // Kilobyte
				faktor = 1024
				strSize = strSize[:len(strSize)-1]
			} else if strSize[len(strSize)-1] == 'M' { // Megabyte
				faktor = 1024 * 1024
				strSize = strSize[:len(strSize)-1]
			}
		}
		size, err := strconv.Atoi(strSize)
		if err != nil {
			fmt.Println(err)
			size = 1
			faktor = 1
		}
		sizeInByte = faktor * size
	}

	delay, err := strconv.Atoi(strDelay)
	if err != nil {
		delay = 0
		if len(strDelay) > 0 {
			fmt.Println(err)
		}
	}

	contentType := httputil.NegotiateContentType(r, []string{"text/plain", "application/octet-stream"}, "text/plain")
	if strRaw == "true" {
		contentType = "application/octet-stream"
	}

	fmt.Println("Got request with delay: " + strconv.Itoa(delay) + ", size: " + strconv.Itoa(sizeInByte) + " and content-type: " + contentType)

	time.Sleep(time.Duration(delay) * time.Millisecond)

	rw.Header()["Content-Type"] = []string{contentType}
	rw.Header()["Content-Length"] = []string{strconv.Itoa(sizeInByte)}

	buffer := make([]byte, buffersize)
	contentTypeHandler := contenttypes.GetContentTypeImplementation(contentType)
	restSizeInByte := sizeInByte
	for restSizeInByte > 0 {
		if restSizeInByte < buffersize {
			buffer = buffer[:restSizeInByte]
		}
		contentTypeHandler.FillBuffer(buffer)
		restSizeInByte -= len(buffer)
		rw.Write(buffer)
	}
}

func main() {

	if len(os.Args) > 2 {
		fmt.Println("Nadire expects zero or one arguments")
		os.Exit(1)
	}

	port := "4321"
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	http.HandleFunc("/", handler)

	fmt.Println("starting Nadire on " + port)
	fmt.Println(http.ListenAndServe(":"+port, nil))
}
