package confy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type DemoStruct struct {
	Port        string `json:"port,omitempty"`
	Address     string `json:"address,omitempty"`
	SessionName string `json:"sessionName,omitempty"`
}

func TestOpen(t *testing.T) {
	// Create json file with empty options
	d := DemoStruct{
		Port:        "",
		Address:     "",
		SessionName: "",
	}
	// convert to bytes
	val, err := json.Marshal(d)
	if err != nil {
		t.Error(err)
	}
	// create temp file
	tmpfile, err := ioutil.TempFile(".", "example.json")
	if err != nil {
		t.Error(err)
	}
	// defer remove
	defer os.Remove(tmpfile.Name())
	// write the json to temp file
	_, err = tmpfile.Write(val)
	if err != nil {
		t.Error(err)
	}
	// close when done writing
	err = tmpfile.Close()
	if err != nil {
		t.Error(err)
	}
	// open file to see how it handles it
	err = Open(tmpfile.Name())
	if err != nil {
		t.Error(err)
	}

	// create json file with correct settings / options
	d2 := DemoStruct{
		Port:        "7000",
		Address:     "nisse",
		SessionName: "nisse123",
	}
	// convert to bytes
	val2, err := json.Marshal(d2)
	if err != nil {
		t.Error(err)
	}
	// create temp file
	tmpfile2, err := ioutil.TempFile(".", "example2.json")
	if err != nil {
		t.Error(err)
	}
	// defer remove
	defer os.Remove(tmpfile2.Name())
	// write to the temp file
	_, err = tmpfile2.Write(val2)
	if err != nil {
		t.Error(err)
	}
	// close when done writing
	err = tmpfile2.Close()
	if err != nil {
		t.Error(err)
	}
	// open to se how it handels it
	err = Open(tmpfile2.Name())
	if err != nil {
		t.Error(err)
	}

	// create json with bad input
	content3 := []byte("{empty}")
	// create temp file
	tmpfile3, err := ioutil.TempFile(".", "example3.json")
	if err != nil {
		t.Error(err)
	}
	// defer remove
	defer os.Remove(tmpfile3.Name())
	// write the content to temp file
	_, err = tmpfile3.Write(content3)
	if err != nil {
		t.Error(err)
	}
	// close when done writing
	err = tmpfile3.Close()
	if err != nil {
		t.Error(err)
	}
	// open file to see how it handles it
	err = Open(tmpfile3.Name())
	if err == nil {
		t.Error("should be error cause empty or wrong options / settings")
	} else {
		fmt.Println(err)
	}

	// open a file that does not exit
	err = Open("DoesNotExist.json")
	if err == nil {
		t.Error("should be error cause empty or wrong options / settings")
	} else {
		fmt.Println(err)
	}
}
