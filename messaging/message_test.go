package messaging

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

type TestArgsData struct {
	Token string      `json:"token"`
	Topic string      `json:"topic"`
	Title string      `json:"title"`
	Body  string      `json:"body"`
	Icon  string      `json:"icon"`
	Data  interface{} `json:"data"`
}

var serverKey = os.Getenv("FIREBASE_SERVER_KEY")

var _ = Describe("Firebase cloud messaging without server key", func() {

	testmessage := TestArgsData{
		Token: "eNstKTcV5Dg:APA91bEGvEHaP6-UdcLBfgaib1lOPUZgrP1QYDAOUoZc_ZQNNlGO1afiR5lGYqbuJTc4YQ0yn3Xogjuj1GeryvvgkcutItfu0kjMwCTIN2CNdp9oiQBPm2394FxHjWMyW8ZgsL1p4xHo",
		Title: "Test cases",
		Body:  "Hello body from cli"}

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(testmessage)
	if errr != nil {
		log.Fatal(errr)
	}

	req, err := http.NewRequest("POST", "/send-message-by-token", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessageByToken)

	handler.ServeHTTP(recorder, req)

	Describe("Send message by token", func() {
		Context("SendMessageByToken", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Firebase cloud messaging negative testing without enviroment variables", func() {

	testmessage := TestArgsData{
		Token: "eNstKTcV5Dg:APA91bEGvEHaP6-UdcLBfgaasasassaswOUoZc_ZQNNlGO1afiR5lGYqbuJTc4YQ0yn3Xogjuj1GeryvvgkcutItfu0kjMwCTIN2CNdp9oiQBPm2394FxHjWMyW8ZgsL1p4xHo",
		Topic: "news",
		Title: "Test cases",
		Body:  "Hello body from cli"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(testmessage)
	if errr != nil {
		log.Fatal(errr)
	}

	req, err := http.NewRequest("POST", "/send-message-by-topic", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessageByTopic)
	handler.ServeHTTP(recorder, req)

	Describe("Send message by topic", func() {
		Context("SendMessageByTopic", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Firebase cloud messaging negative testing for token", func() {

	testmessage := []byte(`{"status":false}`)

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(testmessage)
	if errr != nil {
		log.Fatal(errr)
	}

	req, err := http.NewRequest("POST", "/send-message-by-token", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessageByToken)

	handler.ServeHTTP(recorder, req)

	Describe("Send message by token", func() {
		Context("SendMessageByToken", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Firebase cloud messaging by token", func() {

	os.Setenv("SERVER_KEY", serverKey)

	testmessage := TestArgsData{
		Token: "eNstKTcV5Dg:APA91bEGvEHaP6-UdcLBfgaib1lOPUZgrP1QYDAOUoZc_ZQNNlGO1afiR5lGYqbuJTc4YQ0yn3Xogjuj1GeryvvgkcutItfu0kjMwCTIN2CNdp9oiQBPm2394FxHjWMyW8ZgsL1p4xHo",
		Title: "Test cases",
		Body:  "Hello body from cli"}

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(testmessage)
	if errr != nil {
		log.Fatal(errr)
	}

	req, err := http.NewRequest("POST", "/send-message-by-token", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessageByToken)

	handler.ServeHTTP(recorder, req)

	Describe("Send message by token", func() {
		Context("SendMessageByToken", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Firebase cloud messaging negative testing with args", func() {

	os.Setenv("SERVER_KEY", serverKey)

	testmessage := []byte(`{"status":false}`)

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(testmessage)
	if errr != nil {
		log.Fatal(errr)
	}

	req, err := http.NewRequest("POST", "/send-message-by-topic", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessageByTopic)
	handler.ServeHTTP(recorder, req)

	Describe("Send message by topic", func() {
		Context("SendMessageByTopic", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Firebase cloud messaging negative testing for topic", func() {

	os.Setenv("SERVER_KEY", serverKey)

	testmessage := TestArgsData{
		Token: "eNstKTcV5Dg:APA91bEGvEHaP6-UdcLBfgaasasassaswOUoZc_ZQNNlGO1afiR5lGYqbuJTc4YQ0yn3Xogjuj1GeryvvgkcutItfu0kjMwCTIN2CNdp9oiQBPm2394FxHjWMyW8ZgsL1p4xHo",
		Topic: "news",
		Title: "Test cases",
		Body:  "Hello body from cli"}

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(testmessage)
	if errr != nil {
		log.Fatal(errr)
	}

	req, err := http.NewRequest("POST", "/send-message-by-topic", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessageByTopic)
	handler.ServeHTTP(recorder, req)

	Describe("Send message by topic", func() {
		Context("SendMessageByTopic", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
