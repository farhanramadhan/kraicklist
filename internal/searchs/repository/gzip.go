package repository

import (
	"bufio"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"challenge.haraj.com.sa/kraicklist/internal/searchs"
)

func (sr *searchRepository) Load(ctx context.Context, filepath string) error {
	// open file
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("unable to open source file due: %v", err)
	}
	defer file.Close()
	// read as gzip
	reader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("unable to initialize gzip reader due: %v", err)
	}
	// read the reader using scanner to contstruct records
	cs := bufio.NewScanner(reader)
	for cs.Scan() {
		var r searchs.Record
		err = json.Unmarshal(cs.Bytes(), &r)
		if err != nil {
			continue
		}
		sr.records = append(sr.records, r)
	}

	return nil
}
func (sr *searchRepository) Search(ctx context.Context, query string) ([]searchs.Record, error) {
	var result []searchs.Record
	for _, record := range sr.records {
		// Make Search Keyword Not Case Sensitive
		if strings.Contains(strings.ToLower(record.Title), strings.ToLower(query)) || strings.Contains(strings.ToLower(record.Content), strings.ToLower(query)) {
			result = append(result, record)
		}
	}
	return result, nil
}
