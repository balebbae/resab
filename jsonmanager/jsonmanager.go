package jsonmanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"os"

	"github.com/balebbae/resaB/models"
)

type JsonManager struct {
	InputJsonPath string
	OutputJsonPath string
}

func (jm *JsonManager) JsonManager() (map[string]models.Employee, error) {
	jsonFile, err := os.Open(jm.InputJsonPath)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to open json file")
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to read json file")
	}

	var employees map[string]models.Employee

	err = json.Unmarshal(byteValue, &employees)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to unmarshal json")
	}
	jsonFile.Close()
	return employees, nil
}

func New(intputJsonPath string, outputJsonPath string) *JsonManager {
	return &JsonManager{
		InputJsonPath: intputJsonPath,
		OutputJsonPath: outputJsonPath,
	}
}