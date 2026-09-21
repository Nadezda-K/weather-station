package main

import (
	"bufio"
    "fmt"
	"os"
//	"strconv"
	"strings"
)

type WeatherData struct {
    airTemp string
    airPressure string
    precipitation string
    windSpeed string
    windDirection string
    humidity string
    dewPoint string
    soilMoisture string
    cloudCover string
}

func GetMessage(reader *bufio.Reader) string {
    msg, _ := reader.ReadString('\n')
	msg = strings.TrimSpace(msg)

	return msg
 }

func FillWeather(s string, data *WeatherData) {
	values := strings.Split(s, ",")
	switch values[0] {
		case "1":
			data.airTemp = values[1]
		case "2":
			data.airPressure = values[1]
		case "7":
			data.precipitation = values[1]
		case "11":
			data.windSpeed = values[1]
		case "12":
			data.windDirection = values[1]
		case "13":
			data.humidity = values[1]
		case "14":
			data.dewPoint = values[1]
		case "15":
			data.soilMoisture = values[1]
		case "22":
			data.cloudCover = values[1]
	}
}

 func ClearData(data *WeatherData) {
	data.airTemp = "NULL"
    data.airPressure = "NULL"
    data.precipitation = "NULL"
   	data.windSpeed = "NULL"
    data.windDirection = "NULL"
    data.humidity = "NULL"
    data.dewPoint = "NULL"
    data.soilMoisture = "NULL"
    data.cloudCover = "NULL"

}


 func OutputData(data *WeatherData){
	fmt.Println("airTemp:" + data.airTemp)
    fmt.Println("airPressure:" + data.airPressure)
    fmt.Println("precipitation:" + data.precipitation)
    fmt.Println("windSpeed:" + data.windSpeed)
    fmt.Println("windDirection:" + data.windDirection)
    fmt.Println("humidity:" + data.humidity)
    fmt.Println("dewPoint:" + data.dewPoint)
    fmt.Println("soilMoisture:" + data.soilMoisture)
    fmt.Println("cloudCover:" + data.cloudCover)
}


func main() {
	reader := bufio.NewReader(os.Stdin)

	data := WeatherData{
    	airTemp : "NULL",
    	airPressure : "NULL",
    	precipitation : "NULL",
   		windSpeed : "NULL",
    	windDirection : "NULL",
    	humidity : "NULL",
    	dewPoint : "NULL",
    	soilMoisture : "NULL",
    	cloudCover : "NULL",
	}

	fmt.Println("--- Weather Station ---")
	for {
		message := GetMessage(reader)

		switch message {
		case "get":
			
			OutputData(&data)
		case "clear":
			ClearData(&data)
		case "exit":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			FillWeather(message, &data)
		}
	}

	//fmt.Println(data)
	
} 
