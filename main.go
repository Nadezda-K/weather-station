package main

import (
	"bufio"
    "fmt"
	"os"
	"strconv"
	"strings"
)

type WeatherData struct {
    airTemp *float64
    airPressure *float64
    precipitation *float64
    windSpeed *float64
    windDirection *float64
    humidity *float64
    dewPoint *float64
    soilMoisture *float64
    cloudCover *float64
}

func GetMessage(reader *bufio.Reader) string {
    msg, _ := reader.ReadString('\n')
	msg = strings.TrimSpace(msg)

	return msg
 }

func FillWeather(s string, data *WeatherData) {
	values := strings.Split(s, ",")
	var value float64
	if values[1] == "NULL" {
		value = nil
	} else {
		value, _ = strconv.ParseFloat(values[1],64)
	}
	switch values[0] {
		case "1":
			data.airTemp = &value
		case "2":
			data.airPressure= &value
		case "7":
			data.precipitation = &value
		case "11":
			data.windSpeed = &value
		case "12":
			data.windDirection = &value
		case "13":
			data.humidity = &value
		case "14":
			data.dewPoint = &value
		case "15":
			data.soilMoisture = &value
		case "22":
			data.cloudCover = &value
	}
}

 func ClearData(data *WeatherData) {
	data.airTemp = nil
    data.airPressure = nil
    data.precipitation = nil
   	data.windSpeed = nil
    data.windDirection = nil
    data.humidity = nil
    data.dewPoint = nil
    data.soilMoisture = nil
    data.cloudCover = nil

}


 func OutputData(data *WeatherData){
	printValue := func (key string, value *float64) {
		if value == nil {
			fmt.Printf("%s:NULL\n", key)
		} else {
			fmt.Printf("%s:%g\n", key, *value)
		}

	}
	printValue("airTemp", data.airTemp)
    printValue("airPressure", data.airPressure)
    printValue("precipitation", data.precipitation)
    printValue("windSpeed", data.windSpeed)
    printValue("windDirection", data.windDirection)
    printValue("humidity", data.humidity)
    printValue("dewPoint", data.dewPoint)
    printValue("soilMoisture", data.soilMoisture)
    printValue("cloudCover", data.cloudCover)
}




func main() {
	reader := bufio.NewReader(os.Stdin)

	data := WeatherData{
    	airTemp : nil,
    	airPressure : nil,
    	precipitation : nil,
   		windSpeed : nil,
    	windDirection : nil,
    	humidity : nil,
    	dewPoint : nil,
    	soilMoisture : nil,
    	cloudCover : nil,
	}

	fmt.Printf("--- Weather Station ---\n")
	for {
		message := GetMessage(reader)

		switch message {
		case "get":
			OutputData(&data)
		case "clear":
			ClearData(&data)
		case "exit":
			fmt.Printf("Exiting...\n")
			os.Exit(0)
		default:
			FillWeather(message, &data)
		}
	}

	//fmt.Println(data)
	
} 
