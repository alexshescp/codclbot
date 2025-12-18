package config

import (
    "encoding/json"
    "fmt"
    "os"
)

func loadFile[T any](path string, target *T) error {
    data, err := os.ReadFile(path)
    if err != nil {
        return fmt.Errorf("failed to read %s: %w", path, err)
    }

    if err := json.Unmarshal(data, target); err != nil {
        return fmt.Errorf("failed to unmarshal %s: %w", path, err)
    }

    return nil
}
