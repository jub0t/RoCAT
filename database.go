package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Storage Data
type Storage struct {
	Path string
	Data []Record
}

// This format is used for the data in the disk
type DatabaseStructure struct {
	Data []Record
}

func New(path string) (Storage, error) {
	dummy, _ := json.Marshal(DatabaseStructure{Data: []Record{}})

	if bytes, err := os.ReadFile(path); err != nil {
		if err := os.WriteFile(path, []byte(dummy), os.ModePerm); err != nil {
			fmt.Println(err)
			return Storage{
				Path: path,
				Data: []Record{},
			}, err
		} else {
			return Storage{
				Path: path,
				Data: []Record{},
			}, nil
		}
	} else {
		var json_data DatabaseStructure
		if err := json.Unmarshal(bytes, &json_data); err != nil {
			if err := os.WriteFile(path, []byte(dummy), os.ModePerm); err != nil {
				fmt.Println(err)
				return Storage{
					Path: path,
					Data: []Record{},
				}, err
			} else {
				return Storage{
					Path: path,
					Data: []Record{},
				}, nil
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
func (x *Storage) GetAll() []Record {
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
func (x *Storage) RecordExists(id int) bool {
	for i := 0; i < len(x.Data); i++ {
		if x.Data[i].Id == id {
			return true
		}
	}

	return false
}

// Take the data from memory and store it into disk
func SaveToDisk(data *Storage) error {
	if bytes, err := json.Marshal(data); err != nil {
		return err
	} else {
		if err := os.WriteFile(data.Path, bytes, os.ModePerm); err != nil {
			return err
		} else {
			return nil
		}
	}
}
