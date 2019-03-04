package messaging

import (
	"bytes"
	"encoding/json"
	"github.com/heaptracetechnology/microservice-firebase/result"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

type TestArgsData struct {
	Token string      `json:"token"`
	Title string      `json:"title"`
	Body  string      `json:"body"`
	Icon  string      `json:"icon"`
	Data  interface{} `json:"data"`
}

var _ = Describe("Firebase cloud messaging", func() {

	testmessage := TestArgsData{}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(testmessage)

	req, err := http.NewRequest("POST", "/send-message-by-token", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()

	Describe("Send message by token", func() {
		Context("SendMessageByToken", func() {
			It("Should result http.StatusOK", func() {
				SendMessageByToken(recorder, req)
				Expect(result.GetResult()).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Firebase cloud messaging", func() {

	testmessage := TestArgsData{}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(testmessage)

	req, err := http.NewRequest("POST", "/send-message-by-topic", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()

	Describe("Send message by topic", func() {
		Context("SendMessageByTopic", func() {
			It("Should result http.StatusOK", func() {
				SendMessageByTopic(recorder, req)
				Expect(result.GetResult()).To(Equal(http.StatusOK))
			})
		})
	})
})
