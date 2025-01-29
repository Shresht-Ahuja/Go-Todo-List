package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)
func saveTodosToFile(todos Todos, filename string) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("error in todos: %w", err)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}