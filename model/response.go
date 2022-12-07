package model

import "time"

//common response
type Response struct {
	Code int    `form:"code" json:"code"`
	Msg  string `form:"msg" json:"msg"`
	//Data interface{} `form:"data" json:"data"` //根据要求传入具体的响应
}

//content the list of urls,use for get function
type GetResponse struct {
	Response
	Urls []UrlInfo `form:"urls" json:"urls"`
}

//return id of user
type CreateResponse struct {
	Response
	Id string `form:"id"`
}

//
type QueryResponse struct {
	Response
	Url UrlInfo `form:"url" json:"url"`
}

//use for record get
type LoginRecord struct {
	Time   time.Time
	Ip     string
	Status string
}

//只需返回基本信息，可以适当添加额外信息
type UpdateResponse struct {
	Response
}
type DeleteResponse struct {
	Response
}
type PauseResponse struct {
	Response
}
