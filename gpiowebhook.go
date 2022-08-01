package main

import "log"
import "io"
import "net/http"
import "os"
import "os/exec"
import "time"

func main() {
    KEY := os.Getenv("KEY")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        key := r.URL.Query().Get("key")
        if (key != KEY) {
            log.Printf("Bad key: %s", key)
            http.Error(w, "", 403)
            return
        }

        go func() {
            log.Println("Triggering pin 11")
            cmd := exec.Command("mraa-gpio", "set", "11", "1")
            if err := cmd.Run(); err != nil {
                log.Fatal(err)
            }

            time.Sleep(time.Millisecond * 100)

            cmd = exec.Command("mraa-gpio", "set", "11", "0")
            if err := cmd.Run(); err != nil {
                log.Fatal(err)
            }
        }()

        io.WriteString(w, "OK!")
    })

    http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "all good")
    })

    log.Println("Listening on localhost:28384...")
    log.Fatal(http.ListenAndServe("localhost:28384", nil))
}
