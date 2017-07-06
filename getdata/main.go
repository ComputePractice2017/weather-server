package main
import (
	"log"
	"net/http"
	"archive/zip"
	"io"
	"os"
	"time"
	"path/filepath"
	"github.com/ComputePractice2017/weather-server/model"
	"bufio"
	"encoding/csv"
	"strconv"
)

import (
	r "gopkg.in/gorethink/gorethink.v3"
	"fmt"
	"gopkg.in/gorethink/gorethink.v3/types"
)
//The main function
func main() {
	InitSesson()
	CreateDBIfNotExist()
	CreateTableIfNotExist()
	log.Print("Hello World. I will can to get file")
	//Download file
	for {
	//err := downloadFile("./RU_Hydrometcentre_42.ASCII_ZIP.zip","http://djbelyak.ru/share/RU_Hydrometcentre_42.ASCII_ZIP.zip")
	err := downloadFile("./RU_Hydrometcentre_42.ASCII_ZIP.zip","http://djbelyak.ru/share/Hydra.zip")
	if err != nil {
		panic(err)
	}
	log.Println("Downloading is good")
	err = unzip("./RU_Hydrometcentre_42.ASCII_ZIP.zip","./RU_Hydrometcentre_42.ASCII_ZIP")
	if err != nil {
		log.Println(err)
	}
	newWeatherHandler("./RU_Hydrometcentre_42.ASCII_ZIP/Hydra.txt")
	/*
	newWeatherHandler("./RU_Hydrometcentre_42.ASCII_ZIP/RU_Hydrometcentre_42_1.txt")
	newWeatherHandler("./RU_Hydrometcentre_42.ASCII_ZIP/RU_Hydrometcentre_42_2.txt")
	newWeatherHandler("./RU_Hydrometcentre_42.ASCII_ZIP/RU_Hydrometcentre_42_3.txt")
	newWeatherHandler("./RU_Hydrometcentre_42.ASCII_ZIP/RU_Hydrometcentre_42_4.txt")
	newWeatherHandler("./RU_Hydrometcentre_42.ASCII_ZIP/RU_Hydrometcentre_42_5.txt")
	newWeatherHandler("./RU_Hydrometcentre_42.ASCII_ZIP/RU_Hydrometcentre_42_6.txt")
	newWeatherHandler("./RU_Hydrometcentre_42.ASCII_ZIP/RU_Hydrometcentre_42_7.txt")
	newWeatherHandler("./RU_Hydrometcentre_42.ASCII_ZIP/RU_Hydrometcentre_42_8.txt")
	newWeatherHandler("./RU_Hydrometcentre_42.ASCII_ZIP/RU_Hydrometcentre_42_9.txt")
	newWeatherHandler("./RU_Hydrometcentre_42.ASCII_ZIP/RU_Hydrometcentre_42_10.txt")
	newWeatherHandler("./RU_Hydrometcentre_42.ASCII_ZIP/RU_Hydrometcentre_42_11.txt")
	*/
	time.Sleep(24 * time.Hour)
	log.Println("Unpacking is good")
}
}

var session *r.Session
func InitSesson() error {
	dbaddress := os.Getenv("RETHINKDB_HOST")
	if dbaddress == "" {
		dbaddress = "localhost"
	}

	log.Printf("RETHINKDB_HOST: %s\n", dbaddress)
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address: dbaddress,
	})
	if err != nil {
		return err
	}

	err = CreateDBIfNotExist()
	if err != nil {
		return err
	}

	err = CreateTableIfNotExist()

	return err
}
func CreateDBIfNotExist() error {
	res, err := r.DBList().Run(session)
	if err != nil {
		return err
	}

	var dbList []string
	err = res.All(&dbList)
	if err != nil {
		return err
	}

	for _, item := range dbList {
		if item == "weather" {
			return nil
		}
	}

	_, err = r.DBCreate("weather").Run(session)
	if err != nil {
		return err
	}

	return nil
}

func CreateTableIfNotExist() error {
	res, err := r.DB("weather").TableList().Run(session)
	if err != nil {
		return err
	}

	var tableList []string
	err = res.All(&tableList)
	if err != nil {
		return err
	}

	for _, item := range tableList {
		if item == "weather" {
			return nil
		}
	}

	_, err = r.DB("weather").TableCreate("weather", r.TableCreateOpts{PrimaryKey: "ID"}).Run(session)
	r.DB("weather").Table("weather").IndexCreate("Location", r.IndexCreateOpts{Geo: true}).Run(session)

	return err
}
//Download the file from site
func downloadFile(filepath string, url string) (err error) {
  // Create the file
  out, err := os.Create(filepath)
  if err != nil  {
    return err
  }
  defer out.Close()
 log.Println("done")
  // Get the data
  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  // Writer the body to file
  _, err = io.Copy(out, resp.Body) 
  
  if err != nil  {
    return err
  }

  return nil
}
//Unzip the file from target
func unzip(archive, target string) error {
	log.Println(archive)
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}
	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}
	
		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}
//Записываем файлы в базу
func newWeatherHandler(filepath string) {
    fmt.Println("Запись в базу началась")
	csvFile, _ := os.Open(filepath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'
	var weat model.WeatherData
	fmt.Println("Я тут")
	
	for {		
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		log.Println(line, len(line))
		mask, err := strconv.ParseInt(line[2], 10, 64)
			if err != nil {
			log.Println(err)
		}
		latitude, err := strconv.ParseFloat(line[3], 64)
		if err != nil {
			log.Println(err)
		}
		longitude, err := strconv.ParseFloat(line[4], 64)
		windzonal, err := strconv.ParseFloat(line[5], 64)
		windmeridional, err := strconv.ParseFloat(line[6], 64)
		atmosphericpressure, err := strconv.ParseInt(line[7], 10, 64)
		humidity, err := strconv.ParseInt(line[8], 10, 64)
		rainfall, err := strconv.ParseFloat(line[9], 64)
		temperaturesurface, err := strconv.ParseFloat(line[10], 64)
		airtemperature, err := strconv.ParseFloat(line[11], 64)
				
		weat =  model.WeatherData{
			DateBegin:           line[0],
			Date:                line[1],
			Mask:                mask,
			Location: 			types.Point{Lat: latitude, Lon: longitude},
			//Latitude: 		     latitude,
			//Longitude: 			 longitude,
			WindZonal:           windzonal,
			WindMeridional:      windmeridional,
			AtmosphericPressure: atmosphericpressure,
			Humidity:            humidity,
			Rainfall:            rainfall,
			TemperatureSurface:  temperaturesurface,
			AirTemperature:      airtemperature,
		}
		weat, err = newWeatherData(weat)
		if err != nil {
			//w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		
	}
	
	//
	
	//Стандартный пакет для скачивания файлов 
	//Потом разархивировать unzip(examples)
	//strconf
	//Разархивировать один раз в сутки
	//Два приложения
	//Первое приложение загрузку из файла в базу данных
	//NewHandle

	//передовать параметры запроса
	
	
	}
 //Этот код из другой функции
 func newWeatherData(w model.WeatherData) (model.WeatherData, error){
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost",
	})
	res, err := r.UUID().Run(session)
	if err != nil {
		return w, err
	}

	var UUID string

	err = res.One(&UUID)
	if err != nil {
		return w, err
	}
	
		w.ID = UUID
	

	res, err = r.DB("weather").Table("weather").Insert(w).Run(session)
	if err != nil {
		return w, err
	}
	return w, nil
}