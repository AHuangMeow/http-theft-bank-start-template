package main

import (
	"github.com/AHuangMeow/hacker-support/httptool"
)

func main() {
	request, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"https://gtainmuxi.muxixyz.com/api/v1/organization/code",
		"",
		httptool.DEFAULT,
	)
	if err != nil {
		panic(err)
	}

	response, err := request.SendRequest()
	if err != nil {
		panic(err)
	}

	passport, err := response.GetHeader("Passport")
	if err != nil {
		panic(err)
	}

	request, err = httptool.NewRequest(
		httptool.GETMETHOD,
		"https://gtainmuxi.muxixyz.com/api/v1/organization/secret_key",
		"",
		httptool.DEFAULT,
	)
	if err != nil {
		panic(err)
	}

	request.SetHeader("Passport", passport[0])

	response, err = request.SendRequest()
	if err != nil {
		panic(err)
	}

	response.ShowBody()
}

