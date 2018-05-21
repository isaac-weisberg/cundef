package cundef

import (
	"encoding/json"
	"io/ioutil"
)

// Read is a shortcut function to read and parse projects
func Read(path string) (*Project, error) {
	var error error
	var data []byte
	data, error = ioutil.ReadFile(path)
	if error != nil {
		return nil, error
	}
	var project = &Project{}
	error = json.Unmarshal(data, project)
	if error != nil {
		return nil, error
	}
	return project, nil
}
