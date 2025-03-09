package csvcontroller

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

var csvPath string

func init() {
	csvPath = "/home/" + os.Getenv("USER") + "/tasks.csv"

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

/**
 * AddTask writes a new task to the CSV file
 * @param description string
 */
func AddTask(description string) {
	if len(description) == 0 {
		fmt.Fprintf(os.Stderr, "Task description is empty\n")
		return
	}

	file, err := os.Open(csvPath)
	if err != nil {
		return
	}
	defer file.Close()

	// Read CSV file
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Get the last ID
	var newId int
	if len(data) > 1 {
		newId, err = strconv.Atoi(data[len(data)-1][0])
		if err != nil {
			panic(err)
		}
		newId++
	}

	// Create a new task with timestamp
	timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	newTask := []string{fmt.Sprintf("%d", newId), description, timestamp, "false"}

	// Append the new task to the CSV file
	file, err = os.OpenFile(csvPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(newTask)

	if err != nil {
		panic(err)
	}

	fmt.Println("Task added successfully")
}

func Complete(taskId int) error {
	// Ouvrir le fichier CSV
	file, err := os.Open(csvPath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	// Lire le fichier CSV
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}

	// Rechercher la tâche par ID
	var selectedTask []string
	var selectedIndex int
	for index, row := range data {
		if id, err := strconv.Atoi(row[0]); err == nil && id == taskId {
			selectedTask = row
			selectedIndex = index
			break
		}
	}

	// Vérifier si la tâche a été trouvée
	if selectedTask == nil {
		return fmt.Errorf("task #%d not found", taskId)
	}

	// Vérifier si la tâche est déjà complétée
	if selectedTask[3] == "true" {
		return fmt.Errorf("task #%d is already completed", taskId)
	}

	// Marquer la tâche comme complétée
	data[selectedIndex][3] = "true"

	// Écrire les modifications dans le fichier CSV
	file, err = os.Create(csvPath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.WriteAll(data); err != nil {
		return fmt.Errorf("failed to write to CSV file: %w", err)
	}

	fmt.Println("task completed successfully")
	return nil
}
