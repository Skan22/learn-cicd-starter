package auth

import (
	"net/http"
	"testing"
)

func Test_api_key(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey ezfhiaeohioaf&é1365")
	apiKey, err := GetAPIKey(header)
	if err != nil {
		t.Errorf("expected error to be %v , got  : %v", nil, err)
	}
	if apiKey != "ezfhiaeohioaf&é1365_dffez" {
		t.Errorf("expected %s , got : %s", "ezfhiaeohioaf&é1365", apiKey)
	}

}
