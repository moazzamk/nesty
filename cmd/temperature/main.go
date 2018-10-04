package main

import (
	"fmt"
	"github.com/jsgoecke/nest"
	"os"
)

func main() {
	clientId := "1ed03378-3db5-4f9"
	clientSecret := "41am8NvY6qRHJrvAxYazJMoRI"
	authCode := "46T53KPYN7TKEEXY"
	token := "c.Ew91catSF0bjiHCESCcWSsZcLq6EpguiHdjJQwnjaN3OB50rKKUWtTaL090XHt72TgqZ8dFfMFg5pE3tJNA4x6oKn2h9vVM3h0dFyf8wvT31J5S9Yq6Lr9bHhL8l0dqNyCzMsf0Q0bNelFEo"
	state := `STATE`

	client := nest.New(clientId, state, clientSecret, authCode)
	client.Token = token

	devicesChan := make(chan *nest.Devices)
	go func(outChannel chan *nest.Devices) {
		fmt.Println(client.DevicesStream, "ssss\n\n")

		client.DevicesStream(func(devices *nest.Devices, err error) {
			fmt.Println(err, "ERRRR")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("JKKKKK")
			outChannel <- devices
		})
	}(devicesChan)

	for devices := range devicesChan {
		thermostat := devices.Thermostats["1YTR60d_kAYwfVlXmTXCRJo84hmSNAcl"]
		if thermostat.HvacMode == "eco" {
			fmt.Println("Mode is ECO")
			continue
		} else {
			fmt.Println("ECHO IS OFF")
		}

		fmt.Println("TEMPERATURE IS:", thermostat.TargetTemperatureF)

		thermostat.SetTargetTempF(77)
	}
}