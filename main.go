package main

import (
	request "Alert/src"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main () {
	medi := request.GetHTTPResponse("https://api.gopanacea.org/gov/proposals")
	// msg := request.formatProposalMessage()
	// fmt.Println(medi.Content)

	// fmt.Println(msg)

	mediID := os.Setenv("CheckProposalID", medi.ID)
	if mediID != nil {
		fmt.Println("err mediID", mediID);
	}
	mediStatus := strconv.Itoa(medi.Status)
	errStatus := os.Setenv("CheckProposalStatus", mediStatus)
	if errStatus != nil {
		fmt.Println("err errStatus", errStatus);
	}

	// fmt.Println("프로포절 아이디 제대로 들어오나요",os.Getenv("CheckProposalID"), os.Getenv("CheckProposalStatus"))

	xpla := request.GetHTTPResponse("https://dimension-lcd.xpla.dev/gov/proposals")
	xplaID := os.Setenv("XplaProposalID", xpla.ID)
	if xplaID != nil {
		fmt.Println("err xplaID, xplaID")
	}
	strXpla := strconv.Itoa(xpla.Status)
	xplaStatus := os.Setenv("XplaStatus", strXpla)
	if xplaStatus != nil {
		fmt.Println("err xplaStatus", xplaStatus)
	}

	// fmt.Println("프로포절 아이디 제대로 들어오나요",os.Getenv("XplaProposalID"), os.Getenv("XplaStatus"))
	ticker := time.NewTicker(12 * time.Hour)
	// ticker := time.NewTicker(2 * time.Second)
	fmt.Printf("%s BOT is running.\n", time.Now().Format("2006-01-02 15:04:05"))
	defer ticker.Stop()

	for range ticker.C{
		envProp := os.Getenv("CheckProposalID")
		xplaProp := os.Getenv("XplaProposalID")
		mediMsg := request.MediProposalMessage(medi)
		xplaMsg := request.XplaProposalMessage((xpla))
		currTime := time.Now().Format("01-02 15:04:05")
		
		// xplaMsg := formatProposalMessage(xpla)
		// fmt.Println("메디아이디, 프로포절, 스테이트", medi.ID, envProp, mediStatus)

		if medi.ID != envProp && mediStatus == "2" {
			request.SendSlackMsg(mediMsg)
			os.Setenv("CheckProposalID", medi.ID)
		} else {
			fmt.Printf("%s MediBloc govBot is Working.\n",currTime )
		}
		if xpla.ID != xplaProp && strXpla == "2" {
			request.SendSlackMsg(xplaMsg)
			os.Setenv("XplaProposalID", xpla.ID)
		} else {
			fmt.Printf("%s Xpla govBot is Working.\n",currTime )
		}

		// switch {
		// case medi.ID == envProp && mediStatus == "3":
		// 	request.SendSlackMsg(mediMsg)
		// 	os.Setenv("CheckProposalID", medi.ID)
		// case xpla.ID == xplaProp && strXpla == "2":
		// 		request.SendSlackMsg(xplaMsg)
		// 		os.Setenv("XplaProposalID", xpla.ID)
		// default:
		// 	fmt.Printf("%s Xpla govBot is Working.\n",currTime )
		// }
	}
}