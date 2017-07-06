package api

import (
	//"encoding/csv"
	//"fmt"
	//"io"
	//"io/ioutil"
	"net/http"
	//"github.com/gorilla/mux"
	//"bufio"
	"encoding/json"
	"log"
	//"os"
	"strconv"
	"github.com/ComputePractice2017/weather-server/model"
	"github.com/gorilla/mux"
	"fmt"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World!")
}

func getAllWeatherDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	lat, err := strconv.ParseFloat(vars["lat"], 64)
	if err != nil {
		log.Println("1")
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	long, err := strconv.ParseFloat(vars["long"], 64)
	if err != nil {
		log.Println("1")
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	data, err := model.GetData(lat, long)
	if err != nil {
		log.Println("1")
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var newdata []model.WeatherData
	for _, item := range (data) {
		newdata = append(newdata, item.Doc)
	}

	if err = json.NewEncoder(w).Encode(newdata); err != nil {
		log.Println("1")
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	
	w.WriteHeader(http.StatusOK)
}
/*
func newWeatherHandler(w http.ResponseWriter, r *http.Request) {
	//var weather model.WeatherData
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	csvFile, _ := os.Open("RU_Hydrometcentre_42_1.txt")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var weat []model.WeatherData
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		weat = append(weat, model.WeatherData{
			ID:                  line[0],
			DateBegin:           line[1],
			Date:                line[2],
			Mask:                line[3],
			Latitude:            line[4],
			Longitude:           line[5],
			WindZonal:           line[6],
			WindMeridional:      line[7],
			AtmosphericPressure: line[8],
			Humidity:            line[9],
			Rainfall:            line[10],
			TemperatureSurface:  line[11],
			AirTemperature:      line[12],
		})
	}
	//
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err := json.Unmarshal(body, &weat); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	//Стандартный пакет для скачивания файлов 
	//Потом разархивировать unzip(examples)
	//strconf
	//Разархивировать один раз в сутки
	//Два приложения
	//Первое приложение загрузку из файла в базу данных
	//NewHandle

	//передовать параметры запроса
	if err := json.Unmarshal(body, &weat); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	weat, err = model.NewWeatherData(weat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(weat); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}
*/