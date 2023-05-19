package database

import (
	"encoding/json"
	"fmt"
	"os"
	"rocat/structs"
)

// Storage Data
type Storage struct {
	Path string
	Data []structs.Record
}

// This format is used for the data in the disk
type DatabaseStructure struct {
	Data []structs.Record
}

// Create new database instance and initialize the database
func New(path string) (Storage, error) {
	var storage Storage = Storage{
		Path: path,
		Data: []structs.Record{},
	}

	dummy, _ := json.Marshal(DatabaseStructure{Data: []structs.Record{}})
	if bytes, err := os.ReadFile(path); err != nil {
		if err := os.WriteFile(path, []byte(dummy), os.ModePerm); err != nil {
			fmt.Println(err)
			return storage, err
		} else {
			return storage, nil
		}
	} else {
		var json_data DatabaseStructure
		if err := json.Unmarshal(bytes, &json_data); err != nil {
			if err := os.WriteFile(path, []byte(dummy), os.ModePerm); err != nil {
				fmt.Println(err)
				return storage, err
			} else {
				return storage, nil
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
func (x *Storage) SaveRecord(record structs.Record) {
	x.Data = append(x.Data, record)

	if err := SaveToDisk(x); err != nil {
		fmt.Println(err)
	}
}

// Get all records in the database
func (x *Storage) GetAll() []structs.Record {
	return x.Data
}

// Get Record By Id
func (x *Storage) GetRecord(id int) structs.Record {
	for i := 0; i < len(x.Data); i++ {
		record := x.Data[i]

		if record.Id == id {
			return record
		}
	}

	return structs.Record{}
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
