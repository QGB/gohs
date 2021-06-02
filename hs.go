package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/google/shlex" //go get github.com/google/shlex 
	"net/http"
	"os/exec"
)

func eval_cmd(w http.ResponseWriter, q *http.Request) {
	app := q.URL.Path[1:]
	tokens, _ := shlex.Split(app)
	fmt.Fprintf(w, "%#v\n", tokens)
	cmd := exec.Command(tokens[0], tokens[1:]...)

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(w, "%#v\n", err)
	}
	fmt.Fprintf(w, errb.String()+"\n"+outb.String()+"\n")

}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	listen := flag.String("l", ":3000", "http listen [address]:port")
	flag.Parse()

	http.HandleFunc("/", eval_cmd)
	http.HandleFunc("/headers", headers)

	fmt.Println("HTTP Server eval_cmd listening at", *listen)
	if err := http.ListenAndServe(*listen, nil); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}

}
