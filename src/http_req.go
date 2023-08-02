package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Content struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ContentWrapper struct {

	Value *Content `json:"value"`
	Title string `json:"title"`
}

type Proposal struct {
	ID      string         `json:"id"`
	Content ContentWrapper `json:"content"`
	Status  int            `json:"status"`
	StartTime string       `json:"voting_start_time"`
	EndTime string         `json:"voting_end_time"`
}

type ProposalResponse struct {
	Result []Proposal `json:"result"`
}

func GetHTTPResponse(url string) Proposal {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Could not send the req : %v", err)
		// 이상한거 보냈을때 슬랙 개인메세지 보내기 추가
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("could not read the resp Body : %v", err)
	}

	var proposalResponse ProposalResponse
	if err := json.Unmarshal(body, &proposalResponse); err != nil {
		log.Fatalf("Couldn't parse JSON: %v", err)
	}
	
	var lastProposal = proposalResponse.Result[len(proposalResponse.Result)-1]
	if lastProposal.Content.Value != nil {
		// fmt.Printf("밸류필드가 있습니다.  %v \n", lastProposal.Content.Value.Title)
		lastProposal.Content.Title = lastProposal.Content.Value.Title
	}

	return lastProposal
}