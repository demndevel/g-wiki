package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

// Load loads properties from propPath.
func Load(propPath string) (map[string]string, error) {
	prop := make(map[string]string)
	file, err := os.Open(propPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	count := 0
	reader := bufio.NewReader(file)
	for {
		count++

		linebytes, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return prop, nil
			}
			return nil, err
		}
		if isPrefix {
			err = errors.New("contains too long line at line " + strconv.Itoa(count))
			return nil, err
		}

		// check line
		line := string(linebytes)
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		equalNum := strings.Count(line, "=")
		if equalNum != 1 {
			err = errors.New("invalid format at line " + strconv.Itoa(count))
			return nil, err
		}

		// check key and value
		equalIndex := strings.Index(line, "=")
		key := strings.TrimSpace(string(line[0:equalIndex]))
		value := strings.TrimSpace(string(line[equalIndex+1:]))
		if key == "" || value == "" {
			err = errors.New("invalid format at line " + strconv.Itoa(count))
			return nil, err
		}

		// check duplicate key
		if _, ok := prop[key]; ok {
			err = errors.New("contains duplicate key: " + key)
			return nil, err
		}

		prop[key] = value
	}
	return prop, nil
}
