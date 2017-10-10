package input

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/atrox/homedir"
)

// RequiredValidator ensures the input is not empty
func RequiredValidator(input string) (interface{}, error) {
	if input == "" {
		return nil, fmt.Errorf("This option is required")
	}

	return input, nil
}

// PathValidator ensures the input is valid looking path
// Returns modified extended path
func PathValidator(input string) (interface{}, error) {
	if input == "" {
		return "", nil
	}

	path, err := homedir.Expand(input)
	if err != nil {
		return nil, err
	}
	return filepath.Clean(path), nil
}

// DirectoryValidator ensures the input is a valid and **existing** directory
// Returns modified extended path
func DirectoryValidator(input string) (interface{}, error) {
	if input == "" {
		return "", nil
	}

	path, err := PathValidator(input)
	if err != nil {
		return nil, err
	}
	pathStr := path.(string)

	info, err := os.Stat(pathStr)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", pathStr)
	}

	return pathStr, nil
}

// FileValidator ensures the input is a valid and **existing** file
// Returns modified extended path
func FileValidator(input string) (interface{}, error) {
	if input == "" {
		return "", nil
	}

	path, err := PathValidator(input)
	if err != nil {
		return nil, err
	}
	pathStr := path.(string)

	info, err := os.Stat(pathStr)
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		return nil, fmt.Errorf("%s is a directory - need a file", pathStr)
	}

	return pathStr, nil
}

// IntegerValidator converts the input to a integer with strconv.Atoi
// If input is empty we return nil instead of an integer so you can set a default on your side
func IntegerValidator(input string) (interface{}, error) {
	if input == "" {
		return nil, nil
	}

	i, err := strconv.Atoi(input)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// BooleanValidator converts the input to a boolean
// If input is empty we return nil instead of an integer so you can set a default on your side
func BooleanValidator(input string) (interface{}, error) {
	if input == "" {
		return nil, nil
	}

	b, err := parseBool(input)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// parseBool returns the boolean value represented by the string.
// Based on strconv.ParseBool but with y, yes, n, no added.
func parseBool(str string) (bool, error) {
	switch strings.ToLower(str) {
	case "1", "t", "true", "y", "yes":
		return true, nil
	case "0", "f", "false", "n", "no":
		return false, nil
	}
	return false, fmt.Errorf("%s is not valid. Please choose between (y)es or (n)o", str)
}
