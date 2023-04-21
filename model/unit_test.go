package model

import "testing"

func TestCreateOne(t *testing.T) {

	var w Word
	w.EnUs = "2"
	w.ZhCn = "1"
	w.CreateOne()
}
