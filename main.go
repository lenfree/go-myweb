package main

import (
        "github.com/kabukky/httpscerts"
        "io"
        "log"
        "net/http"
        "os"
)

func returnHostname(w http.ResponseWriter, r *http.Request) {
        name, err := os.Hostname()
        if err != nil {
                io.WriteString(w, err.Error())
        } else {
                io.WriteString(w, name)
        }
}

func redirect(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w,
            r,
            "https://" + r.Host + r.URL.String(),
            301,
        )
}

func main() {
        // Generate SSL cert if doesn't exists
        // Thanks to https://github.com/kabukky/httpscerts
        err := httpscerts.Check("cert.pem", "key.pem")
        if err != nil {
                // Generate new SSL cert
                err := httpscerts.Generate("cert.pem", "key.pem", "localhost")
                if err != nil {
                        log.Fatal("Error generating new cert")
                }
        }

        http.HandleFunc("/", returnHostname)
        go http.ListenAndServe(":80", http.HandlerFunc(redirect))
        http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
}
