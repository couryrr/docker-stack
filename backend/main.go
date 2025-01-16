package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"log/slog"
	"net/http"
)

func main(){
    mux := http.NewServeMux()

    mux.HandleFunc("POST /enc", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        defer r.Body.Close()
        message, err := io.ReadAll(r.Body)
        if err != nil {
            slog.Error("bad things", "err", err)
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        h := sha256.New()
        h.Write(message)
        hash := h.Sum(nil)
        enc := hex.EncodeToString(hash)
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte(enc))
    })

    log.Fatal(http.ListenAndServe(":8080", mux))
}
