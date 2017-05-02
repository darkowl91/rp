package rp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// decodeError decodes an Error from an io.Reader.
func decodeError(r io.Reader) error {
	var e struct {
		Code    int    `json:"error_code, omitempty"`
		Message string `json:"message"`
	}
	err := json.NewDecoder(r).Decode(&e)
	if err != nil {
		log.Error(err)
		return errors.New("couldn't decode RP error")
	}
	return fmt.Errorf("Code: %d, Mesage: %s", e.Code, e.Message)
}
