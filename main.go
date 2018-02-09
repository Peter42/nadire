package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"math/rand"
	"os"
)


const BUFFERSIZE = 1024 * 1024

const randomStringContent = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.:"; // length of 64
func getRandomString(buffer []byte, rng *rand.Rand){
	for i := range buffer {
        buffer[i] = randomStringContent[rng.Int() & 63]
    }
}

func handler(rw http.ResponseWriter, r *http.Request) {
	strDelay := r.URL.Query().Get("delay")
	strSize  := r.URL.Query().Get("size")
	
	sizeInByte := 42
	if(len(strSize) > 0) {
		faktor := 1
		if(len(strSize) > 1) {
			if(strSize[len(strSize)-1] == 'k') { // Kilobyte
				faktor = 1024
				strSize = strSize[:len(strSize)-1]
			} else if(strSize[len(strSize)-1] == 'M') { // Megabyte
				faktor = 1024 * 1024
				strSize = strSize[:len(strSize)-1]
			}
		}
		size, err := strconv.Atoi(strSize)
		if (err != nil) {
			fmt.Println(err)
			size = 1
			faktor = 1
		}
		sizeInByte = faktor * size
	}

	delay, err := strconv.Atoi(strDelay)
	if (err != nil) {
		delay = 0
		if (len(strDelay) > 0) {
			fmt.Println(err)
		}
	}
	
	
	fmt.Println("Got request with delay: " + strconv.Itoa(delay) + " and size: " + strconv.Itoa(sizeInByte))
	
	time.Sleep(time.Duration(delay) * time.Millisecond)
	
	rw.Header()["Content-Type"] = []string{"text/plain"}
	rw.Header()["Content-Length"] = []string{strconv.Itoa(sizeInByte)}
	
    rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	buffer := make([]byte, BUFFERSIZE)
	restSizeInByte := sizeInByte
	for (restSizeInByte > 0) {
		if(restSizeInByte < BUFFERSIZE) {
			buffer = buffer[:restSizeInByte]
		}
		getRandomString(buffer, rng)
		restSizeInByte -= len(buffer)
		rw.Write(buffer)
	}
}

func main() {

	if(len(os.Args) > 2) {
		fmt.Println("Nadire expects zero or one arguments")
		os.Exit(1)
	}
	
	port := "4321"
	if(len(os.Args) == 2) {
		port = os.Args[1]
	}

	http.HandleFunc("/", handler)

	fmt.Println("starting Nadire on " + port)
	fmt.Println(http.ListenAndServe(":" + port, nil))
}