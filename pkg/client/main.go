package client

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func CallServer() {
	response, err := http.Get("http://0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		log.Info("response:", string(body))
	} else {
		log.Fatal("Error! Status:", response.StatusCode)
	}
}
