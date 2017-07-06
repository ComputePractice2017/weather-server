package model

import (
	r "gopkg.in/gorethink/gorethink.v3"
	  "gopkg.in/gorethink/gorethink.v3/types"
	"log"
	)

//WeatherData is a struct to store date of weather
type WeatherData struct {
	ID        string `json:"id",gorethink:"id"`
	DateBegin string `json:"dateBegin",gorethink:"DateBegin`
	Date      string `json:"date",gorethink:"Date `
	Mask      int64  `json:"mask",gorethink:"mask"`
	Location  types.Point  `json:"location",gorethink:"location"`
	//Latitude            float64 `json:"latitude",gorethink:"latitude`
	//Longitude           float64 `json:"longitude",gorethink:"longitude`
	WindZonal           float64 `json:"windZonal",gorethink:"windZonal`
	WindMeridional      float64 `json:"windMeridional",gorethink:"windMeridional`
	AtmosphericPressure int64   `json:"atmosphericPressure",gorethink:"atmosphericPressure`
	Humidity            int64   `json:"humidity,gorethink:"humidity`
	Rainfall            float64 `json:"rainfall",gorethink:"rainfall`
	TemperatureSurface  float64 `json:"temperatureSurface",gorethink:"temperatureSurface`
	AirTemperature      float64 `json:"airTemperature",gorethink:"airTemperature"`
}

type Wrapper struct {
	Dist float64
	Doc WeatherData
}

var session *r.Session

func InitSession() error {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address: "localhost",
	})
	return err
}

//GetData is a simple storage of weathers
func GetData(lat, long float64) ([]Wrapper, error) {
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost",
	})

	if err != nil {
		log.Println("1")
		return nil, err
	}

	var Base = r.Point(lat, long)
	log.Println(lat, long)
	//Зпрос на выборку приближенных значений из базы данных
	res, err := r.DB("weather").Table("weather").GetNearest(Base, r.GetNearestOpts{Index: "Location"}).Run(session)
	
	//res, err := r.DB("weather").Table("weather").Run(session)
	if err != nil {
		log.Println("1")
		return nil, err
	}

	var response []Wrapper
	err = res.All(&response)
	if err != nil {
		log.Println("1")
		return nil, err
	}
	log.Println(response)
	return response, nil

	// Output:
	// Hello World
}
