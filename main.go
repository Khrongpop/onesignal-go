package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tbalthazar/onesignal-go"
)

var appID = os.Getenv("OSAPPID")
var appKey = os.Getenv("OSAPIKEY")
var client = onesignal.NewClient(nil)

func main() {
	client.UserKey = appID
	client.AppKey = appKey

	http.HandleFunc("/", index)
	http.HandleFunc("/push", push)
	port := os.Getenv("PORT")
	fmt.Println(`start port http:localhost:` + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func index(w http.ResponseWriter, r *http.Request) {

}

func push(w http.ResponseWriter, r *http.Request) {
	// playerID := "db480b54-6734-4957-9be8-2a7b8d2b0f20"
	notificationReq := &onesignal.NotificationRequest{
		AppID:    appID,
		Contents: map[string]string{"en": "English message"},
		IsIOS:    true,
		IsChrome: true,
		// IncludePlayerIDs: []string{playerID},
		IncludedSegments: []string{"All"},
	}
	NotificationCreateResponse, res, err := client.Notifications.Create(notificationReq)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Print(`NotificationCreateResponse : `)
	fmt.Println(NotificationCreateResponse)
	fmt.Print(`*NotificationCreateResponse : `)
	fmt.Println(*NotificationCreateResponse)
	fmt.Print(`&NotificationCreateResponse : `)
	fmt.Println(&NotificationCreateResponse)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response Body:", string(body))
}
