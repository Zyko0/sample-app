package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func worker() {
	for {
		fmt.Println("Dlrow olleH")
		time.Sleep(100000000)
	}
}

func main() {
	var firstArg = flag.Bool("worker", false, "random help message")
	flag.Parse()
	if *firstArg == false {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello World\n")
			fmt.Println("myenvvar:" + os.Getenv("MYENVVAR") + "\n")
		})
		address := fmt.Sprintf(":%s", os.Getenv("PORT"))
		log.Fatal(http.ListenAndServe(address, nil))
	} else {
		worker()
	}
}
