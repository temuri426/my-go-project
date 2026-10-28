package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)

func main() {

	cmd := exec.Command("cat", "/etc/os-release")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, string(stdoutStderr))

	}

	http.HandleFunc("/api", h1)
	http.ListenAndServe(":8090", nil)
}

