package main

import "github.com/stianeikeland/go-rpio/v4"
import "log"
import "io"
import "net/http"
import "os"
import "time"

func main() {
    KEY := os.Getenv("KEY")

    err := rpio.Open()
    if err != nil {
        log.Fatalln("Could not initialize rpio: ", err)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        key := r.URL.Query().Get("key")
        if (key != KEY) {
            log.Printf("Bad key: %s", key)
            http.Error(w, "", 403)
            return
        }

        go func() {
            log.Println("Triggering pin 17")
            pin := rpio.Pin(17)
            pin.Output()
            pin.High()
            time.Sleep(time.Millisecond * 100)
            pin.Low()
        }()

        io.WriteString(w, "OK!")
    })

    log.Println("Listening on localhost:28384...")
    log.Fatal(http.ListenAndServe("localhost:28384", nil))
}
