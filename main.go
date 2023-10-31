package main

import (
	"context"
	"elasticsearch-learn/pkg/elastic"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	certificate, err := os.ReadFile("http_ca.crt")
	if err != nil {
		log.Fatal("unable to read http certificate because of ", err)
	}
	ec := elastic.NewElasticClient(elasticsearch.Config{
		Addresses: []string{"https://127.0.0.1:9200"},
		Username:  "elastic",
		Password:  "<YOUR_PASSWORD>",
		CACert:    certificate,
	}, elastic.ElasticDependencies{})

	ec.SearchEngineInfo(context.Background())
}
