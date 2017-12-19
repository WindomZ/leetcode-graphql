package leetcodegraphql

type Response struct {
	IsCurrentUserAuthenticated bool         `json:"isCurrentUserAuthenticated"`
	Data                       ResponseData `json:"data"`
}

type ResponseData struct {
	Question interface{} `json:"question"`
}
