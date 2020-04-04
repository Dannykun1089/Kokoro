//
//Functions relating to opening, reading, and writing files
//

package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//FileExists - Checks if a file exists, returns false if the path is a directory or is not there
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//PathExists - Checks if the given path exists, returns false if the path dosen't lead to anything
func PathExists(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsNotExist(err)
}

//JSONFileToStruct - Reads the specified file into the passed in struct pointer
func JSONFileToStruct(path string, struc interface{}) error {
	//Open the file
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	//Read the file's bytes
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	//Unmarshal the JSON bytes into the struct
	err = json.Unmarshal(data, struc)
	if err != nil {
		return err
	}
	return nil
}
