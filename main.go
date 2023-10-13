package main

import (
	request "Alert/src"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main () {
	ticker := time.NewTicker(3 * time.Second)
	// ticker := time.NewTicker(30 * time.Minute)
	// ticker := time.NewTicker(2 * time.Second)
	fmt.Printf("%s BOT is running.\n", time.Now().Format("2006-01-02 15:04:05"))
	initMedi := request.GetHTTPResponse("https://api.gopanacea.org/gov/proposals")
	initXpla := request.GetHTTPResponse("https://dimension-lcd.xpla.dev/gov/proposals")
	os.Setenv("initMediProposalID", initMedi.ID)
	os.Setenv("initXplaProposalID", initXpla.ID)
	initMediID := os.Getenv("initMediProposalID")
	initXplaID := os.Getenv("initXplaProposalID")
	fmt.Printf("Init ProposalID of Medibloc %s \n",initMediID)
	fmt.Printf("Init ProposalID of XPLA %s \n",initXplaID)
	defer ticker.Stop()

	for range ticker.C{
		// HTTP.GET
		CurrMedi := request.GetHTTPResponse("https://api.gopanacea.org/gov/proposals")
		CurrXpla := request.GetHTTPResponse("https://dimension-lcd.xpla.dev/gov/proposals")

		// medibloc Prop.ID 환경변수 등록
		// fmt.Println(medi.ID)
		os.Setenv("MediProposalID", CurrMedi.ID)
		
		// medibloc status string 변환
		mediStatus := strconv.Itoa(CurrMedi.Status)
		// medibloc Status 환경변수 등록
		// fmt.Println(mediStatus)
		os.Setenv("CheckProposalStatus", mediStatus)


		os.Setenv("XplaProposalID", CurrXpla.ID)
		strXpla := strconv.Itoa(CurrXpla.Status)
		os.Setenv("XplaStatus", strXpla)
		
		// PrevMediID := os.Getenv("MediProposalID")
		// PrevXplaID := os.Getenv("XplaProposalID")
		mediMsg := request.MediProposalMessage(CurrMedi)
		xplaMsg := request.XplaProposalMessage(CurrXpla)
		currTime := time.Now().Format("01-02 15:04:05")
		
		fmt.Printf("PrevMediID:%v, CurrMediID:%v, PrevXplaID:%v , CurrXplaID:%v \n",initMediID,CurrMedi.ID, initXplaID, CurrXpla.ID)
		fmt.Printf("mediStatus: %v", mediStatus)
		if CurrMedi.ID != initMediID && mediStatus == "2" {
			request.SendSlackMsg(mediMsg)
			initMediID = CurrMedi.ID
			fmt.Printf("[New Proposal on Medibloc] #%v -> #%v\n", initMediID, CurrMedi.ID)
		} else {
			fmt.Printf("[%s] MediBloc govBot is Working.\n",currTime )
		}
		if CurrXpla.ID != initXplaID && strXpla == "2" {
			request.SendSlackMsg(xplaMsg)
			initXplaID = CurrXpla.ID
			fmt.Printf("[New Proposal on XPLA] #%v -> #%v\n", initXplaID, CurrXpla.ID)
		} else {
			fmt.Printf("[%s] XPLA govBot is Working.\n",currTime )
		}
	}
}