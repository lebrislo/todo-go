package csvcontroller

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

var csvPath string

func init() {
	initCsv()
}

func initCsv() {
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
func AddTask(description string) error {
	if len(description) == 0 {
		return fmt.Errorf("Task description is empty")
	}

	file, err := os.Open(csvPath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	// Read CSV file
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}

	// Get the last ID
	var newId int
	if len(data) > 1 {
		newId, err = strconv.Atoi(data[len(data)-1][0])
		if err != nil {
			return err
		}
		newId++
	}

	// Create a new task with timestamp
	timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	newTask := []string{fmt.Sprintf("%d", newId), description, timestamp, "false"}

	// Append the new task to the CSV file
	file, err = os.OpenFile(csvPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(newTask)

	if err != nil {
		return fmt.Errorf("failed to write CSV file: %w", err)
	}

	return nil
}

func CompleteTask(taskId int) error {
	file, err := os.Open(csvPath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

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

func ListTasks(allTask bool) error {
	// Create a new tabwriter.Writer instance.
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer w.Flush()

	// Open csv task file
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
	// Remove the header row
	if len(data) > 0 {
		data = data[1:]
	}

	var tasks [][]string
	for _, task := range data {
		if allTask {
			tasks = append(tasks, task)
		} else if task[3] == "false" {
			tasks = append(tasks, task)
		}
	}

	// Write some data to the Writer.
	if allTask {
		fmt.Fprintln(w, "ID\tTask\tCreated\tDone")
	} else {
		fmt.Fprintln(w, "ID\tTask\tCreated")
	}

	for _, task := range tasks {
		var timestamp, _ = strconv.Atoi(task[2])
		var time = time.Unix(int64(timestamp), 0)
		var timeStr = timediff.TimeDiff(time)
		if allTask {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", task[0], task[1], timeStr, task[3])
		} else {
			fmt.Fprintf(w, "%s\t%s\t%s\n", task[0], task[1], timeStr)
		}
	}

	return nil
}

func DeleteTask(taskId int) error {
	file, err := os.Open(csvPath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}

	var taskFound = false
	for index, row := range data {
		if id, err := strconv.Atoi(row[0]); err == nil && id == taskId {
			data = append(data[:index], data[index+1:]...)
			taskFound = true
			break
		}
	}

	if !taskFound {
		return fmt.Errorf("task #%d not found", taskId)
	}

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

	fmt.Println("task deleted successfully")
	return nil
}

func DeleteAll() error {
	err := os.Remove(csvPath)
	if err != nil {
		return fmt.Errorf("failed to remove CSV file %w", err)
	}

	initCsv()

	return nil
}
