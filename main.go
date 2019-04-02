package main

import (
	"fmt"
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
	// listOpt := &onesignal.NotificationListOptions{
	// 	AppID:  appID,
	// 	Limit:  10,
	// 	Offset: 0,
	// }
	// _, _, _ := client.Notifications.List(listOpt)

	// playerID := "db480b54-6734-4957-9be8-2a7b8d2b0f20"
	// notificationReq := &onesignal.NotificationRequest{
	// 	AppID:    appID,
	// 	Contents: map[string]string{"en": "English message"},
	// 	IsIOS:    true,
	// 	IsChrome: true,
	// 	// IncludePlayerIDs: []string{playerID},
	// 	IncludedSegments: []string{"All"},
	// }
	// _, res, err := client.Notifications.Create(notificationReq)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res)
	// body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println("response Body:", string(body))

	http.HandleFunc("/", index)
	port := os.Getenv("PORT")
	fmt.Println(`start port http:localhost:` + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func index(w http.ResponseWriter, r *http.Request) {
	client := onesignal.NewClient(nil)
	client.UserKey = appID
	client.AppKey = appKey
	// appRequest := &onesignal.AppRequest{
	// 	Name: "Your app 1",
	// }
	// app, res, err := client.Apps.Create(appRequest)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res)
	// fmt.Println(app)

}
