package main

import (
	"github.com/AHuangMeow/hacker-support/httptool"
)

func main() {
	passport := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiQUh1YW5nTWVvdyIsImlhdCI6MTc2Mjk5NDkwMywibmJmIjoxNzYyOTk0OTAzfQ.t3uPobRQKbBE0xqafESebOCSXNfk4gvG8bHKXQfnSro"

	request, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"https://gtainmuxi.muxixyz.com/api/v1/muxi/backend/computer/examination",
		"",
		httptool.DEFAULT,
	)
	if err != nil {
		panic(err)
	}

	request.SetHeader("Passport", passport)

	response, err := request.SendRequest()
	if err != nil {
		panic(err)
	}

	response.ShowBody()
}
