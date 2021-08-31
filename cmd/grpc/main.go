package run

import (
	"otp_service/initiator"
)

var testInit bool

func init() {
	//flag.BoolVar(&testInit, "test", false, "initialize test mode without serving")
	//flag.Parse()
	//
	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

}

func Run() {
	initiator.InitRest(testInit)
}
