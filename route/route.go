package route

import (
    "github.com/gorilla/mux"
    "github.com/heaptracetechnology/microservice-firebase/messaging"
    "log"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "SendMessageByToken",
        "POST",
        "/sendMessageByToken",
        messaging.SendMessageByToken,
    },
    Route{
        "SendMessageByTopic",
        "POST",
        "/sendMessageByTopic",
        messaging.SendMessageByTopic,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
