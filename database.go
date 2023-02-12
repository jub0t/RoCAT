package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Storage Data
type Storage struct {
	Path string
	Data []Record
}

type DatabaseStructure struct {
	Data []Record
}

func New(path string) (Storage, error) {
	if bytes, err := os.ReadFile(path); err != nil {
		// Failed to read file
		initFiles([]string{path})
		return Storage{}, err
	} else {
		var json_data DatabaseStructure
		if err := json.Unmarshal(bytes, &json_data); err != nil {
			dummy, err := json.Marshal(DatabaseStructure{Data: []Record{}})

			if err != nil {
				fmt.Println("Unable to marshal dummy json data")
				return Storage{}, err
			}

			if err := os.WriteFile(path, []byte(dummy), os.ModePerm); err != nil {
				fmt.Println(err)
				return Storage{}, err
			} else {
				return Storage{}, nil

			}
		} else {
			return Storage{
				Path: path,
				Data: json_data.Data,
			}, nil
		}
	}
}

// Save A Record To Database
func (x *Storage) SaveRecord(record Record) {
	x.Data = append(x.Data, record)

	if err := SaveToDisk(x); err != nil {
		fmt.Println(err)
	}
}

// Get all records in the database
func (x *Storage) GetAll(record Record) []Record {
	return x.Data
}

// Get Record By Id
func (x *Storage) GetRecord(id int) Record {
	for i := 0; i < len(x.Data); i++ {
		record := x.Data[i]

		if record.Id == id {
			return record
		}
	}

	return Record{}
}

// Get all records with the given name
func (x *Storage) SearchRecord(name string) []Record {
	var records []Record

	for i := 0; i < len(x.Data); i++ {
		record := x.Data[i]

		if strings.ToLower(record.Name) == strings.ToLower(name) {
			records = append(records, record)
		}
	}

	return records
}

// Take the data from memory and store it into disk
func SaveToDisk(data *Storage) error {
	if bytes, err := json.Marshal(data.Data); err != nil {
		return err
	} else {
		fmt.Println(data.Path)
		if err := os.WriteFile(data.Path, bytes, os.ModePerm); err != nil {
			return err
		} else {
			return nil
		}
	}
}
