package main

import (
    "log"
    "net/http"
    "os"
    "sync/atomic"
    "strconv"
)

func main() {
    hostname, _ := os.Hostname()
    var counter int64
    var max int64

    maxi := os.Getenv("MAX_TIME")
    if len(maxi) == 0 {
        max = 5
    } else {
        maxint,err := strconv.ParseInt(maxi, 10, 32)
        if err!= nil {
          log.Printf("Bad value for MAX_TIME")
          os.Exit(1)
        }
        max = maxint
    }
   
    log.Printf("I will die after %d get",max)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Received request from ", r.RemoteAddr)
        atomic.AddInt64(&counter, 1)

        if counter > max {
           log.Printf("I can't survive more than %d \n",max)
           os.Exit(1)
        }

        if counter > (max-1) {
            w.WriteHeader(http.StatusInternalServerError)
            log.Printf("I will die after the next get... \n")
            w.Write([]byte("I'm not well. Please restart me!\n"))
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("You've hit " + hostname + " - " + strconv.FormatInt(counter, 10) +" \n"))
    })

    log.Print("Kubia server starting...")
    http.ListenAndServe(":8080", nil)
}
