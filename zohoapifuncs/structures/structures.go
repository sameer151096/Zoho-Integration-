package structures

import "time"

// can also be acquired using postman and then put in configuration file for further usage. It is simpler in my opinion.
type FirstResponse_AcReToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresInSec int    `json:"expires_in_sec"`
	APIDomain    string `json:"api_domain"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

type RefreshedAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresInSec int    `json:"expires_in_sec"`
	APIDomain    string `json:"api_domain"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

// values that are used to build url
type Request struct {
	Grant_type    string `json: "grant_type"`
	Client_id     string `json:"client_id"`
	Client_secret string `json:"client_secret"`
	Redirect_uri  string `json:"redirect_uri"`
	Code          string `json:"code"`
}

//payload for zoho api request
type RequestData struct {
	Company       string `json:"Company"`
	LastName      string `json:"Last_Name"`
	FirstName     string `json:"First_Name"`
	Email         string `json:"Email"`
	State         string `json:"State"`
	Phone         string `json:"Phone"`
	LeadGenerator string
	CreatedBy     string `json:"Created By"`
	Source        string `json:"Source"`
	LeadOwner     string `json:"Lead Owner"`
}

type ZohoApiRequest struct {
	Payload []RequestData `json:"data"`
	Trigger []string
}

type ZohoApiResponse struct {
	ResponseData []struct {
		Code    string `json:"code"`
		Details struct {
			APIName      string    `json:"api_name"`
			ModifiedTime time.Time `json:"Modified_Time"`
			ModifiedBy   struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"Modified_By"`
			CreatedTime time.Time `json:"Created_Time"`
			ID          string    `json:"id"`
			CreatedBy   struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"Created_By"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

// Use case specific Lead definitions
type Lead struct {
	LeadID               string // must be generate each time a lead is created..\
	Email                string `json:"Email"`
	Name                 string `json:"Name"`
	Phone                string `json:"Phone"`
	TimeOfLeadGeneration string
	LeadConverted        bool
}
