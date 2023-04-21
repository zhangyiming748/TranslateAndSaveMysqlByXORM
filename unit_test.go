package main

import (
	"TranslateAndSaveMysqlByXORM/model"
	"fmt"
	"testing"
)

func TestClean(t *testing.T) {
	var w model.Word
	w.DeleteAll()
}

func TestGetAll(t *testing.T) {
	var w model.Word
	words := w.GetAll()
	for _, word := range words {
		fmt.Printf("%+v\n", word)
	}
}
