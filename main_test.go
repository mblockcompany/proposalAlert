package main

import (
	request "Alert/src"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCheckForUpdates 함수는 GetHTTPResponse 함수가 지속적으로 변경되는 응답을 처리할 수 있는지 테스트합니다.
func TestCheckForUpdates(t *testing.T) {
	// Mock server를 생성합니다.
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// 간단한 JSON 응답을 반환합니다. 실제 응답 형식에 맞게 변경해야 합니다.
		rw.Write([]byte(`{"ID": "testID"}`))
	}))
	// 테스트 종료 후 server 종료
	defer server.Close()

	// initMedi 값 설정
	initMedi := request.GetHTTPResponse(server.URL) // mock server URL 사용

	// CurrMedi 값 설정
	CurrMedi := request.GetHTTPResponse(server.URL)

	// initMedi와 CurrMedi가 동일한지 확인
	if CurrMedi.ID != initMedi.ID {
		t.Errorf("Expected ID %v, but got %v", initMedi.ID, CurrMedi.ID)
	}

	// 추가적으로 다른 테스트 케이스나 검증 로직을 추가할 수 있습니다.
}