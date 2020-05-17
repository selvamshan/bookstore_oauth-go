package oauth

import (
	"os"
	"fmt"
	"testing"
	"net/http"
	"strconv"
	"github.com/stretchr/testify/assert"
	"github.com/mercadolibre/golang-restclient/rest"
)


func TestMain(m *testing.M){
	fmt.Println("about to start oauth tests")
	rest.StartMockupServer()
	os.Exit(m.Run())
}


func TestOauthCosntants(t *testing.T){
	assert.EqualValues(t,  "X-Public", headerXPublic)
	assert.EqualValues(t,  "X-Client-Id", headerXClientId)
	assert.EqualValues(t,  "X-Caller-Id", headerXCallerId)
	assert.EqualValues(t,  "access_token", paramAccessToken)
}


func TestIsPublicNilRequest(t *testing.T){
	assert.True(t, IsPublic(nil))
}

func TestIsPublicNoError(t *testing.T){
	request := http.Request{
		Header: make(http.Header),
	}
	assert.False(t, IsPublic(&request))

	request.Header.Add("X-Public", "true")
	assert.True(t, IsPublic(&request))
}


func TestGetCallerIdNilRequest(t *testing.T){
	assert.EqualValues(t, 0, GetCallerId(nil))
}


func TestGetCallerIdInvalidCallerFormat(t *testing.T){
	//TODO:
}


func TestGetCallerIdNoError(t *testing.T){
	request := http.Request{
		Header: make(http.Header),
	}
	request.Header.Add("X-Caller-Id", "7")
	callerId, err := strconv.ParseInt(request.Header.Get(headerXCallerId), 10, 64)
	assert.Nil(t,err)
	assert.EqualValues(t, 7, callerId)
}

func TestGetAccessTokenInvalidResetClientResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodGet,
		URL:          "http://localhost:8080/oauth/access_token/abc123",
		ReqBody: ``, 
		RespHTTPCode: -1,
		RespBody: `{}`,

	})

	accessToken, err := GetAccessToken("abc123")
	assert.Nil(t, accessToken)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "invalid restclient response when trying to get access token", err.Message())
}