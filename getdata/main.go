package main
import (
	"log"
	"net/http"
	"archive/zip"
	"io"
	"os"
	"time"
	"path/filepath"
)



func main() {
	log.Print("Hello World. I will can to get file")
	//Download file
	for {
	downloadFile("./","http://djbelyak.ru/share/RU_Hydrometcentre_42.ASCII_ZIP.zip")
	log.Println("Downloading is good")
	unzip("getdata/RU_Hydrometcentre_42.ASCII_ZIP.zip","/getdata")
	time.Sleep(24 * time.Hour)
	log.Println("Unpacking is good")
	}

}

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