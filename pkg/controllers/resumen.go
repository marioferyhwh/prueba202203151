package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/marioferyhwh/prueba202203151/pkg/models"
)

func GetResumen(fechaInit string, dias int64) models.AnswerA {

	var clients = []models.Client{}
	var i int64
	res1 := strings.Split(fechaInit, "-")

	anio, _ := strconv.ParseInt(res1[0], 10, 64)
	mes, _ := strconv.ParseInt(res1[1], 10, 64)
	dia, _ := strconv.ParseInt(res1[2], 10, 64)
	var fecha = time.Date(int(anio), time.Month(mes), int(dia), 0, 0, 0, 0, time.UTC)

	fmt.Println(fecha)
	for i = 0; i < dias; i++ {
		var fecha1 = fmt.Sprintf("%d-", fecha.Year())
		if int(fecha.Month()) <= 9 {
			fecha1 += "0"
		}
		fecha1 += fmt.Sprintf("%d-", int(fecha.Month()))
		if fecha.Day() <= 9 {
			fecha1 += "0"
		}
		fecha1 += fmt.Sprintf("%d", fecha.Day())
		clientsTemp := getPageToJson("https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/" + fecha1)
		clients = append(clients, clientsTemp...)
		fmt.Println(fecha1)
		fecha = fecha.Add(24 * time.Hour)
	}
	var answer = models.AnswerA{}
	for _, client := range clients {
		if client.Compro {
			answer.Total += client.Monto
			if answer.CompraMasAlta < client.Monto {
				answer.CompraMasAlta = client.Monto
			}
			if client.Tdc != "" {
				answer.ComprasPorTDC.Oro += int(client.Monto)
				answer.ComprasPorTDC.Amex += 1
			}
		} else {
			answer.NoCompraron += 1
		}
	}
	return answer
}

func getPageToJson(baseURL string) []models.Client {
	var clients = []models.Client{}
	resp, getErr := http.Get(baseURL)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	jsonErr := json.Unmarshal(body, &clients)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return clients
}
