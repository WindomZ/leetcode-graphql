package leetcodegraphql

// Response the structure of leetcode response
type Response struct {
	IsCurrentUserAuthenticated bool         `json:"isCurrentUserAuthenticated"`
	Data                       ResponseData `json:"data"`
}

// ResponseData the structure of leetcode response data
type ResponseData struct {
	Question interface{} `json:"question"`
}
