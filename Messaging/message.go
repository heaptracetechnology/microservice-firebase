package Messaging

import (
	"encoding/json"
	"io/ioutil"
	"github.com/NaySoftware/go-fcm"
	"net/http"
	"os"
)

type Message struct {
    Success int `json:"success"`
    TopicMsgid string `json:"topic_msgid"`
	Status string `json:"status"`
}

type MesssageToTopic struct {
    Topic string `json:"topic"`
    Body map[string]string `json:"body"`
}

//Send Message
func SendMessage(w http.ResponseWriter, r *http.Request) {

	var serverKey = os.Getenv("SERVER_KEY")
	
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	messsageToTopic := new(MesssageToTopic)
	err = json.Unmarshal([]byte(body), messsageToTopic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	c := fcm.NewFcmClient(serverKey)
	c.NewFcmMsgTo(messsageToTopic.Topic, messsageToTopic.Body)
	status, err := c.Send()

	bytes, err := json.Marshal(status)
	
	WriteJsonResponse(w, bytes, 200)
}

func WriteJsonResponse(w http.ResponseWriter, bytes []byte, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
