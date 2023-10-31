package main

import (
	"elasticsearch-learn/model"
	"fmt"
	"time"
)

func main() {
	// ec := elasticsearch.NewElasticClient(elasticsearch.ElasticDependencies{

	// })

	// sampleDOB, err := time.Parse(time.DateOnly, "2001-04-25")
	// if err != nil {
	// 	log.Fatal("unable to parse day of birth because of ", err)
	// }

	// err = ec.InsertData(context.Background(), model.User{
	// 	FirstName:  "Joshua",
	// 	MiddleName: "Ryandafres",
	// 	LastName:   "Pangaribuan",
	// 	Age:        22,
	// 	DayOfBirth: sampleDOB,
	// 	Address:    "Jl. Dr. Susilo IIC, No.61, Jakarta Barat",
	// })

	// fmt.Println(ec)

	joshua := model.User{
		FirstName:  "Joshua",
		MiddleName: "Ryandafres",
		LastName:   "Pangaribuan",
		Age:        20,
		DayOfBirth: time.Now(),
		Address:    "Jl. Dr. Susilo IIC, No. 61. Jakarta Barat",
	}

	fmt.Println(joshua.FullName())
}
