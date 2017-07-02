package model

import (
	r "gopkg.in/gorethink/gorethink.v3"
)

//WeatherData is a struct to store date of weather
type WeatherData struct {
	ID                  string `json:"id",gorethink:"id"`
	DateBegin           string `json:"dateBegin",gorethink:"DateBegin`
	Date                string `json:"date",gorethink:"Date `
	Mask                string/*int*/ `json:"mask",gorethink:"mask"`
	Latitude            string/*float32*/ `json:"latitude",gorethink:"latitude`
	Longitude           string/*float32*/ `json:"longitude",gorethink:"longitude`
	WindZonal           string/*float32*/ `json:"windZonal",gorethink:"windZonal`
	WindMeridional      string/*float32*/ `json:"windMeridional",gorethink:"windMeridional`
	AtmosphericPressure string/*int*/ `json:"atmosphericPressure",gorethink:"atmosphericPressure`
	Humidity            string/*int*/ `json:"humidity,gorethink:"humidity`
	Rainfall            string/*float32*/ `json:"rainfall",gorethink:"rainfall`
	TemperatureSurface  string/*float32*/ `json:"temperatureSurface",gorethink:"temperatureSurface`
	AirTemperature      string/*float32*/ `json:"airTemperature",gorethink:"airTemperature"`
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
func GetData() ([]WeatherData, error) {
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost",
	})
	if err != nil {
		return nil, err
	}

	res, err := r.DB("weather").Table("weather").Run(session)
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
func NewWeatherData(w []WeatherData) ([]WeatherData, error) {
	res, err := r.UUID().Run(session)
	if err != nil {
		return w, err
	}

	var UUID string

	err = res.One(&UUID)
	if err != nil {
		return w, err
	}
	for i := 0; i < len(w); i++ {
		w[i].ID = UUID
	}
	res, err = r.DB("weather").Table("weather").Insert(w).Run(session)
	if err != nil {
		return w, err
	}
	//
	return w, nil
}
