package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func GetAccessToken(clientId string, clientSecret string) (string, float64){
	// Request for access token
	accessTokenResponse, error := http.PostForm(ConfigApiEndpoint + "/oauth2/token/", url.Values{
		"grant_type": {"client_credentials"},
		"scope": {"upload_documentation"},
		"client_id": {clientId},
		"client_secret": {clientSecret}})

	if error != nil {
		fmt.Println(error)
		return "", 0
	}

	var accessTokenResult map[string]interface{}

	json.NewDecoder(accessTokenResponse.Body).Decode(&accessTokenResult)

	fmt.Println(accessTokenResult)

	accessToken := accessTokenResult["access_token"].(string)
	expireDate := accessTokenResult["expires_in"].(float64)

	return accessToken, expireDate


}
