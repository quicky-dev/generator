package util

import (
    "os"
    "path"

    "github.com/google/uuid"
)
// CreateFile creates the output script and writes it to a file
func CreateFile(filePath string, script []string) (string, error) {
    // Generate a new uuid4
    uuid, err := uuid.NewRandom(); if err != nil {
        return "", err
    }

    // Convert the uuid4 to a string
    fileName := uuid.String()
    file := path.Join(filePath, fileName)

    // Attempt to create a file
    f, err := os.Create(file); if err != nil {
        return "", err
    }

    // Write the script to file
    for _, command := range script {
        _, err := f.WriteString(command + "\n"); if err != nil {
            f.Close()
            return "", err
        }
    }

    // Success
    f.Close()
    return file, nil
}


