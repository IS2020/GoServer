package main
import (
	"fmt"
	"github.com/NaySoftware/go-fcm"
)
const (
	serverKey = "API-KEY"
	topic = "/topics/all"
)
func sendNotification()  {
	payload := &fcm.NotificationPayload{
		Title:  "Alerta de prueba",
		Body: "Alerta sismica",
	}
	c := fcm.NewFcmClient(serverKey)
	c.SetNotificationPayload(payload)

	c.NewFcmMsgTo(topic, nil)


	status, err := c.Send()


	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}
}
