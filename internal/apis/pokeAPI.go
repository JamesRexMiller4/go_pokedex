package internal

import (
	"fmt"
	"io"
	"net/http"
)

func GetLocation(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET request failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code != 200")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return data, nil
}
