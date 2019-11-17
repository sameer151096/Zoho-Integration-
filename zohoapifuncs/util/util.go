package util

import (
	"Zoho-Integration-/zohoapifuncs/structures"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"
)

// GenerateFirstAcessAndRefreshToken : helper function to generate First access token and refresh tokens
func GenerateFirstAcessAndRefreshToken() structures.FirstResponse_AcReToken {
	u := &url.URL{
		Scheme: "https",
		Host:   "accounts.zoho.in", // careful with the hosts.. they change with the country
		Path:   "oauth/v2/token",
	}
	value := u.Query()

	// values code to generate url for 1st access token ..

	value.Add("grant_type", "authorization_code") // this was "authorization_code"
	value.Add("client_id", "1000.V0GHWQWT5FK682485BGZ8Q4OF72CQV")
	value.Add("client_secret", "bde1e775ca0ef16157c077c3c8eb718fc4ba6208c6")
	value.Add("redirect_uri", "http://www.abc.com")
	value.Add("code", "1000.d0b4d3ed49dbf0704447daae354fcbfd.ddc94d2acbf3953f6acc86195f401dc7")
	// values added and now encode them into Url's struct var RawQuery
	u.RawQuery = value.Encode()

	client := http.Client{Timeout: time.Second * 10}
	//Url formed should be of similar signature - https://accounts.zoho.in/oauth/v2/token?client_id=1000.V0GHWQWT5FK682485BGZ8Q4OF72CQV&client_secret=bde1e775ca0ef16157c077c3c8eb718fc4ba6208c6&code=1000.d20e7ec7cbf609231adb7c2a40c6ca98.990123ddf6fa9c301a2f94c77703fcae&grant_type=authorization_code&redirect_uri=http%3A%2F%2Fwww.abc.com

	response, err := client.Post(u.String(), "application/json", nil)
	if err != nil {
		log.Print("\n leadManager error while generating access token. Error is ", err)
	}

	// this piece of code is for FirstResponse_AcReToken
	resobj := structures.FirstResponse_AcReToken{}
	err = json.NewDecoder(response.Body).Decode(&resobj)
	if err != nil {
		log.Print("\n error while decoding response.error  :", err)
	}
	return resobj
}

// GenerateAcessTokenUsingRefreshToken : helper function to generate access token using refresh token genrated
func GenerateAcessTokenUsingRefreshToken() structures.RefreshedAccessTokenResponse {
	u := &url.URL{
		Scheme: "https",
		Host:   "accounts.zoho.in",
		Path:   "oauth/v2/token",
	}
	value := u.Query()

	// code to generate access token using refresh token
	value.Add("grant_type", "refresh_token")
	value.Add("client_id", "1000.V0GHWQWT5FK682485BGZ8Q4OF72CQV")
	value.Add("client_secret", "bde1e775ca0ef16157c077c3c8eb718fc4ba6208c6")
	value.Add("refresh_token", "1000.bd8fd45b5b9b7163e3e847917bcd6928.511caf66f9f18d11f2c2dea7fa6572b1")

	// values added and now encode them into Url's struct var RawQuery
	u.RawQuery = value.Encode()
	client := http.Client{Timeout: time.Second * 10}
	response, err := client.Post(u.String(), "application/json", nil)
	if err != nil {
		log.Print("\n leadManager error while generating access token. Error is ", err)
	}

	resobj := structures.RefreshedAccessTokenResponse{}
	err = json.NewDecoder(response.Body).Decode(&resobj)
	if err != nil {
		log.Print("\n error while decoding response.error  :", err)
	}
	return resobj
}

// ConvertLeadsManDataToZohoNativeRequest : helper function to convert custom lead definitions to zoho native definitions
func ConvertLeadsManDataToZohoNativeRequest(lead *structures.Lead) structures.RequestData {
	data := structures.RequestData{}
	data.FirstName = lead.Name
	data.LastName = lead.Name // Revisit: do something about lastname.
	data.Email = lead.Email
	data.Phone = lead.Phone
	data.CreatedBy = ""
	data.Source = ""
	data.LeadOwner = ""
	return data
	// add mandatory fields in this function if it's the use case. This function can act as enforcer of the required fields..
}
