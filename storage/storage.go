package storage

import (
    "encoding/json"
    "os"

    "github.com/Homework900/simple-library2/domain"
    "github.com/Homework900/simple-library2/library"
)

type Storable interface {
    Save() error
    Load() error
}

type storageData struct {
    Books   []*domain.Book   `json:"books"`
    Readers []*domain.Reader `json:"readers"`
}

func SaveLibraryToJSON(filePath string, lib *library.Library) error {
    data := storageData{Books: lib.Books, Readers: lib.Readers}
    jsonData, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(filePath, jsonData, 0644)
}

func LoadLibraryFromJSON(filePath string) (*library.Library, error) {
    jsonData, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }
    var data storageData
    if err := json.Unmarshal(jsonData, &data); err != nil {
        return nil, err
    }
    lib := library.New()
    lib.Books = data.Books
    lib.Readers = data.Readers
    lib.UpdateIDs()
    return lib, nil
}