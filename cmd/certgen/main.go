package main

import (
	"certgen/model"
	"certgen/pkg/cert"
	"certgen/pkg/cert/html"
	"certgen/pkg/cert/pdf"

	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	PATH_OUT  = "./out/"
	PDF_TYPE  = "pdf"
	HTML_TYPE = "html"
)

func init() {
	err := createOutDirectory()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Flags to perform specific actions
	typeInput := flag.String("type", "pdf", "pdf or html ")
	pathInput := flag.String("file", "./example.csv", "/path/myfile.csv")

	flag.Parse()

	file, err := os.Open(*pathInput)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	var generator cert.Service

	switch *typeInput {
	case PDF_TYPE:
		generator = pdf.New()
	case HTML_TYPE:
		generator = html.New()
	}

	for {
		line, err := csvReader.Read()
		if line == nil || err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			//todo
		}

		if len(line) < 3 {
			log.Printf("want 3 datas got %d: %v", len(line), line)
			continue
		}

		//skip header
		if line[0] == "enseignement" &&
			line[1] == "name" &&
			line[2] == "date" {
			continue
		}
		data := &model.Cert{
			Education: line[0],
			Name:      line[1],
			Date:      line[2],
		}

		err = generator.GenCert(data)
		if err != nil {
			log.Println(err)
			//todo
		}
		path := PATH_OUT + data.CertName
		err = os.WriteFile(path, data.Certificat, os.ModePerm)
		if err != nil {
			log.Fatal(err)
			//todo
		}
	}

}

func createOutDirectory() error {

	if _, err := os.Stat(PATH_OUT); os.IsNotExist(err) {
		err := os.MkdirAll(PATH_OUT, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error repo not be create : %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("error validation repo : %v", err)
	}
	return nil
}
