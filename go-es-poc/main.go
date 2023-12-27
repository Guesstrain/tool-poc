package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type Student struct {
	Name         string  `json:"name"`
	Age          int64   `json:"age"`
	AverageScore float64 `json:"average_score"`
}

func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err
}

func InsertES() {
	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing :", err)
		panic("Client fail")
	}

	newStudent := Student{
		Name:         "steve Doe",
		Age:          25,
		AverageScore: 85.0,
	}

	dataJSON, err := json.Marshal(newStudent)
	js := string(dataJSON)
	_, err = esclient.Index().Index("students").BodyJson(js).Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful")
}

func QueryES() []Student {
	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing: ", err)
		panic("Client fail")
	}

	var students []Student

	seachSource := elastic.NewSearchSource()
	seachSource.Query(elastic.NewMatchQuery("name", "Doe"))

	queryStr, err1 := seachSource.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	}
	fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))

	searchService := esclient.Search().Index("students").SearchSource(seachSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
	}

	for _, hit := range searchResult.Hits.Hits {
		var student Student
		err := json.Unmarshal(hit.Source, &student)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

		students = append(students, student)
	}
	if err != nil {
		fmt.Println("Fetching student fail: ", err)
		return nil
	} else {
		return students
	}
}

func main() {
	//InsertES()
	students := QueryES()
	for _, s := range students {
		fmt.Printf("Student found Name: %s, Age: %d, Score: %f \n", s.Name, s.Age, s.AverageScore)
	}
}
