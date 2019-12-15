package main

import (
	"fmt"

	"github.com/NaySoftware/go-fcm"
)

const (
	serverKey = "API-KEY"
	topic     = "/topics/all"
)

func sendNotification() {
	payload := &fcm.NotificationPayload{
		Title: "Alerta de prueba",
		Body:  "Alerta sismica",
	}
	data := map[string]interface{}{
		"click_action": "FLUTTER_NOTIFICATION_CLICK",
		"comida":       "Alerta sísmica detectada",
	}
	c := fcm.NewFcmClient(serverKey)
	c.SetNotificationPayload(payload)

	c.NewFcmMsgTo(topic, data)

	status, err := c.Send()

	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}
}
