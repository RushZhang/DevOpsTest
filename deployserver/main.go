package main

import (
	"io"
	"net/http"
	"os/exec"
	"log"
)

func reLaunch() {
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func deployPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "My Deploy Server, ReLaunching...")
	reLaunch() //因为reLaunch花的时间久，是网络，比io久，所以要放在io后边
}

func main() {
	http.HandleFunc("/", deployPage)
	http.ListenAndServe(":6000", nil)
} 