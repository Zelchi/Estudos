package main

import (
	"github.com/zelchi/testego"
)

func main() {
	fruta := testego.CreateFruit("Batata", 20.50)
	testego.Start(fruta)
	fruta.Selling()
	fruta.Marketing()
}
