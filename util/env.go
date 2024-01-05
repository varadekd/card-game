package util

import (
	"fmt"
	"os"
)

func GetEnvVariable(variable string) (string, error) {
	val, variableFound := os.LookupEnv(variable)

	if !variableFound {
		return val, fmt.Errorf("we were unable to find variable %s in the application environment", variable)
	}

	if val == "" {
		return val, fmt.Errorf("no value found for variable %s in the application environment", variable)
	}

	return val, nil
}
