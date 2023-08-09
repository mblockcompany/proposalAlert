package main

import (
	request "Alert/src"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main () {
	// ticker := time.NewTicker(1 * time.Hour)
	// ticker := time.NewTicker(30 * time.Minute)
	ticker := time.NewTicker(2 * time.Second)
	fmt.Printf("%s BOT is running.\n", time.Now().Format("2006-01-02 15:04:05"))
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
		
		PrevMediID := os.Getenv("MediProposalID")
		PrevXplaID := os.Getenv("XplaProposalID")
		mediMsg := request.MediProposalMessage(CurrMedi)
		xplaMsg := request.XplaProposalMessage(CurrXpla)
		currTime := time.Now().Format("01-02 15:04:05")
		
		fmt.Printf("PrevMediID:%v, CurrMediID:%v, PrevXplaID:%v , CurrXplaID:%v \n",PrevMediID,CurrMedi.ID, PrevXplaID, CurrXpla.ID)
		if CurrMedi.ID != PrevMediID && mediStatus == "2" {
			request.SendSlackMsg(mediMsg)
			os.Setenv("MediProposalID", CurrMedi.ID)
			fmt.Printf("[New Proposal on Medibloc] #%v -> #%v\n", PrevMediID, CurrMedi.ID)
		} else {
			fmt.Printf("[%s] MediBloc govBot is Working.\n",currTime )
		}
		if CurrXpla.ID != PrevXplaID && strXpla == "2" {
			request.SendSlackMsg(xplaMsg)
			os.Setenv("XplaProposalID", CurrXpla.ID)
			fmt.Printf("[New Proposal on XPLA] #%v -> #%v\n", PrevXplaID, CurrXpla.ID)
		} else {
			fmt.Printf("[%s] XPLA govBot is Working.\n",currTime )
		}
	}
}