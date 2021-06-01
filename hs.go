package main

import (
    "fmt"
    "bytes"
    "net/http"
    //"regexp"
     "os/exec"
     "github.com/google/shlex" //go install github.com/google/shlex@latest
    
)

func hello(w http.ResponseWriter, q *http.Request) {
   app := q.URL.Path[1:]
   tokens, _ := shlex.Split(app)
   fmt.Fprintf(w,"%#v\n", tokens)
   cmd := exec.Command(tokens[0],tokens[1:]...)

  var outb, errb bytes.Buffer
  cmd.Stdout = &outb
  cmd.Stderr = &errb
  err := cmd.Run()
  if err != nil {
     fmt.Fprintf(w,"%#v\n", err)
     //return
  }
  fmt.Fprintf(w, outb.String()+"\n"+errb.String()+"\n")  
    
    
   /* 
   
  fmt.Println("out:", outb.String(), "err:", errb.String())
    fmt.Fprintf(w, q.URL.Path[1:]+"\n"+string(out)+"\n")  
       out, _ := ioutil.ReadAll(cmd.Stdout)
    

  if (err != nil) {
       fmt.Fprintln(w, err.String())
       return
    } 
   
   
    cmd, err := exec.Run(app, []string{app, "-l"}, nil, "", exec.DevNull, exec.Pipe, exec.Pipe)

    if (err != nil) {
       fmt.Fprintln(w, err.String())
       return
    }

    var b bytes.Buffer
    io.Copy(&b, cmd.Stdout)
    fmt.Println(b.String())

    cmd.Close()

   out, _ := exec.Command(q.URL.Path[1:]).Output()
  
    fmt.Fprintf(w, q.URL.Path[1:]+"\n"+string(out)+"\n")
    */
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
    
    http.HandleFunc("/", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8090", nil)
}
