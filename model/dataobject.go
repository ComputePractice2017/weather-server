package model

import (
	r "gopkg.in/gorethink/gorethink.v3"
)

//WeatherData is a struct to store date of weather
type WeatherData struct {
	ID                  string `json:"id",gorethink:"id"`
	DateBegin           string `json:"dateBegin",gorethink:"DateBegin`
	Date                string `json:"date",gorethink:"Date `
	Mask                int64 `json:"mask",gorethink:"mask"`
	Latitude            float64 `json:"latitude",gorethink:"latitude`
	Longitude           float64 `json:"longitude",gorethink:"longitude`
	WindZonal           float64 `json:"windZonal",gorethink:"windZonal`
	WindMeridional      float64 `json:"windMeridional",gorethink:"windMeridional`
	AtmosphericPressure int64 `json:"atmosphericPressure",gorethink:"atmosphericPressure`
	Humidity            int64 `json:"humidity,gorethink:"humidity`
	Rainfall            float64 `json:"rainfall",gorethink:"rainfall`
	TemperatureSurface  float64 `json:"temperatureSurface",gorethink:"temperatureSurface`
	AirTemperature      float64 `json:"airTemperature",gorethink:"airTemperature"`
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
func GetData(long, lat float64) ([]WeatherData, error) {
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost",
	})

	if err != nil {
		return nil, err
	}

	var Base = r.Point(long, lat);
	res, err := r.DB("weather").GetNearest(Base, r.GetNearestOpts{Index: "location", MaxResults: 1}).Run(session)

	//res, err := r.DB("weather").Table("weather").Run(session)
	if err != nil {
		return nil, err
	}

	var response []WeatherData
	err = res.All(&response)
	if err != nil {
		return nil, err
	}

	return response, nil

	// Output:
	// Hello World
}

