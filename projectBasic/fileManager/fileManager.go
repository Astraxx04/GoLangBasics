package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if(err != nil) {
		return nil, errors.New("error occured while opening the file")
	}

	defer file.Close() // This will be executed when the code/ function is executed

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err() 

	if(err != nil) {
		// file.Close()
		return nil, errors.New("error occured while reading the file")
	}
	// file.Close()
	return lines, err
}

func (fm FileManager) WriteResult(data interface{})(error) {
	file, err := os.Create(fm.OutputFilePath)
	if(err != nil) {
		return errors.New("error in creating the file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if(err != nil) {
		// file.Close()
		return errors.New("error in converting data to JSON")
	}
	// file.Close()
	return nil
}// any can aloso be used

func New(inputPath string, outputPath string) (FileManager) {
	return FileManager {
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}