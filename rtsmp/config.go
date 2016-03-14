package rtsmp

import (
   "github.com/hashicorp/hcl"
)

func loadConfig(path string)(map[string]interface{}, error) {
	var result map[string]interface{}
	err := hcl.Load(path, &result)
	return result, err
}