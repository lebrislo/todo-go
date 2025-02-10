package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const csvPath = "/home/" + os.Getenv("USER") + "/tasks.csv"

func csvInit() {
	file, err := os.Open(csvPath)
	if err != nil {
		file, err = os.Create(csvPath)
		if err != nil {
			panic(err)
		}
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"ID", "Description", "CreatedAt", "IsComplete"}
	writer.Write(headers)
}

func AddTask(description string) {
	if len(description) == 0 {
		fmt.Fprintf(os.Stderr, "Task description is empty")
	}

	file, err := os.Open(csvPath)
	if err != nil {
		return
	}

	// find next task ID
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println("data len", len(data))
}
