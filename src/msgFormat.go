package request

import (
	"fmt"
	"time"
)

func MediProposalMessage(proposal Proposal) string {
	s1 := proposal.StartTime
	e1 := proposal.EndTime
	startTime, _ := time.Parse(time.RFC3339, s1)
	endTime, _ := time.Parse(time.RFC3339, e1)

	return fmt.Sprintf(`[New proposal on MediBloc]
제안 개요 : %v
투표 기간 : %s ~ %s
투표페이지 : https://www.mintscan.io/medibloc/proposals/%s`, proposal.Content.Title, startTime.Format("2006-01-02 15:04"), endTime.Format("2006-01-02 15:04"),proposal.ID)
}

func XplaProposalMessage(proposal Proposal) string {
	s1 := proposal.StartTime
	e1 := proposal.EndTime
	startTime, _ := time.Parse(time.RFC3339, s1)
	endTime, _ := time.Parse(time.RFC3339, e1)

	return fmt.Sprintf(`[New proposal on XPLA]
제안 개요 : %v
투표 기간 : %s ~ %s
투표페이지 : https://www.mintscan.io/xpla/proposals/%s`, proposal.Content.Title, startTime.Format("2006-01-02 15:04"), endTime.Format("2006-01-02 15:04"),proposal.ID)
}