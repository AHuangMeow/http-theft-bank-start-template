package main

import (
	"github.com/AHuangMeow/hacker-support/httptool"
)

func main() {
	request, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/iris_sample",
		"",
		httptool.DEFAULT,
	)
	if err != nil {
		panic(err)
	}

	request.SetHeader("Passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiQUh1YW5nTWVvdyIsImlhdCI6MTc2Mjk5NDkwMywibmJmIjoxNzYyOTk0OTAzfQ.t3uPobRQKbBE0xqafESebOCSXNfk4gvG8bHKXQfnSro")

	response, err := request.SendRequest()
	if err != nil {
		panic(err)
	}

	response.Save("./iris.png")

	request, err = httptool.NewRequest(
		httptool.POSTMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate",
		"./iris.png",
		httptool.FILE,
	)
	if err != nil {
		panic(err)
	}

	request.SetHeader("Passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiQUh1YW5nTWVvdyIsImlhdCI6MTc2Mjk5NDkwMywibmJmIjoxNzYyOTk0OTAzfQ.t3uPobRQKbBE0xqafESebOCSXNfk4gvG8bHKXQfnSro")

	response, err = request.SendRequest()
	if err != nil {
		panic(err)
	}

	response.ShowBody()
}
