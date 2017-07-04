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
)
//The main function
func main() {
	log.Print("Hello World. I will can to get file")
	//Download file
	/*for {
	err := downloadFile("./weatherdata.zip","http://djbelyak.ru/share/RU_Hydrometcentre_42.ASCII_ZIP.zip.")
	if err != nil {
		panic(err)
	}*/
	log.Println("Downloading is good")
	err := unzip("./Hyndra.zip","./")
	if err != nil {
		panic(err)
	}
	newWeatherHandler("./Hyndra/Hyndra.txt")
	//newWeatherHandler(http.ResponseWriter, http.Request);
	time.Sleep(24 * time.Hour)
	log.Println("Unpacking is good")
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

	csvFile, _ := os.Open(filepath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var weat []model.WeatherData
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		mask, err := strconv.ParseInt(line[3], 10, 64)
			if err != nil {
			log.Println(err)
		}
		latitude, err := strconv.ParseFloat(line[4], 64)
		longitude, err := strconv.ParseFloat(line[5], 64)
		windzonal, err := strconv.ParseFloat(line[6], 64)
		windmeridional, err := strconv.ParseFloat(line[7], 64)
		atmosphericpressure, err := strconv.ParseInt(line[8], 10, 64)
		humidity, err := strconv.ParseInt(line[9], 10, 64)
		rainfall, err := strconv.ParseFloat(line[10], 64)
		temperaturesurface, err := strconv.ParseFloat(line[11], 64)
		airtemperature, err := strconv.ParseFloat(line[12], 64)
				
		weat = append(weat, model.WeatherData{
			ID:                  line[0],
			DateBegin:           line[1],
			Date:                line[2],
			Mask:                mask,
			Latitude:            latitude,
			Longitude:           longitude,
			WindZonal:           windzonal,
			WindMeridional:      windmeridional,
			AtmosphericPressure: atmosphericpressure,
			Humidity:            humidity,
			Rainfall:            rainfall,
			TemperatureSurface:  temperaturesurface,
			AirTemperature:      airtemperature,
		})
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
	
	weat, err := newWeatherData(weat)
	if err != nil {
		//w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	}
 //Этот код из другой функции
 func newWeatherData(w []model.WeatherData) ([]model.WeatherData, error){
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost",
	})
	if err != nil {
		return nil, err
	}
	res, err := r.UUID().Run(session)
	if err != nil {
		return w, err
	}

	var UUID string

	err = res.One(&UUID)
	if err != nil {
		return w, err
	}
	for i:=0;i<len(w);i++ {
		w[i].ID = UUID
	res, err = r.DB("weather").Table("weather").Insert(w[i]).Run(session)
}
	
	if err != nil {
		return w, err
	}
	//
	return w, nil
}
 