package util

import (
    "os"
    "path"

    "github.com/google/uuid"
)

// GenerateUUID generates a new UUID4
func GenerateUUID() (string, error) {
    uuid, err := uuid.NewRandom(); if err != nil  {
        return "", err
    }

    return uuid.String(), nil
}
// CreateFile creates the output script and writes it to a file
func CreateFile(filePath string, script []string) (string, error) {

    // Generate a new UUID
    fileName, err := GenerateUUID(); if err != nil {
        return "", err
    }
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


