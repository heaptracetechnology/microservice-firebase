package route

import (
    "github.com/gorilla/mux"
    "github.com/microservice-firebase/Messaging"
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
        "/send-message-by-token",
        Messaging.SendMessageByToken,
    },
    Route{
        "SendMessageByTopic",
        "POST",
        "/send-message-by-topic",
        Messaging.SendMessageByTopic,
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
