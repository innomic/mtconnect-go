package main

import (
	"encoding/xml"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/innomic/mtconnect-go/schema/v2.2/mtstreams"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	resp, err := http.Get("https://mtconnect.trakhound.com/current")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if http.StatusOK != resp.StatusCode {
		return errors.New(resp.Status)
	}
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var unmarshalled mtstreams.MTConnectStreamsType
	if err := xml.Unmarshal(respBodyBytes, &unmarshalled); err != nil {
		return err
	}
	marshalled, err := xml.MarshalIndent(unmarshalled, "", "  ")
	if err != nil {
		return err
	}
	// Use unmarshalled data
	log.Println(string(marshalled))
	return nil
}
