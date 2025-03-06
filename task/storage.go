package task

import (
	"encoding/json"
	"os"
	"sync"
)

const dataFile = "data/tasks.json"

var (
	taskCache   []Task
	cacheLock   sync.Mutex
	cacheLoaded bool
)

// LoadTasks loads tasks from JSON into memory
func LoadTasks() ([]Task, error) {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	// Return cached data if already loaded
	if cacheLoaded {
		return taskCache, nil
	}

	// Read from file
	file, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			taskCache = []Task{}
			cacheLoaded = true
			return taskCache, nil
		}
		return nil, err
	}

	// Unmarshal JSON
	err = json.Unmarshal(file, &taskCache)
	if err != nil {
		return nil, err
	}

	cacheLoaded = true
	return taskCache, nil
}

// SaveTasks writes tasks to JSON file and updates cache
func SaveTasks(tasks []Task) error {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	taskCache = tasks
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(dataFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
