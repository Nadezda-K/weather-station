package main

import (
	"bufio"
    "fmt"
	"os"
	"strconv"
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
	value, err := strconv.ParseFloat(values[1],64)

	var value_str string
	if err != nil {
		value_str = "NULL"
	} else {
		value_str = fmt.Sprintf("%g", value)
	}

	switch values[0] {
		case "1":
			data.airTemp = value_str
		case "2":
			data.airPressure= value_str
		case "7":
			data.precipitation = value_str
		case "11":
			data.windSpeed = value_str
		case "12":
			data.windDirection = value_str
		case "13":
			data.humidity = value_str
		case "14":
			data.dewPoint = value_str
		case "15":
			data.soilMoisture = value_str
		case "22":
			data.cloudCover = value_str
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
	fmt.Printf("airTemp:%v\n", data.airTemp)
    fmt.Printf("airPressure:%v\n", data.airPressure)
    fmt.Printf("precipitation:%v\n", data.precipitation)
    fmt.Printf("windSpeed:%v\n", data.windSpeed)
    fmt.Printf("windDirection:%v\n", data.windDirection)
    fmt.Printf("humidity:%v\n", data.humidity)
    fmt.Printf("dewPoint:%v\n", data.dewPoint)
    fmt.Printf("soilMoisture:%v\n", data.soilMoisture)
    fmt.Printf("cloudCover:%v\n", data.cloudCover)
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
