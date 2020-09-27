package types

import "time"

// ServiceTypeProduct ..
const (
	ServiceTypeProduct int = 2 + iota
	ServiceTypOrder
	ServiceTypLogistics
)

//Service ..
type Service struct {
	ServiceID   string    `json:"service_id"`
	ServiceName string    `json:"service_name"`
	ExpireTime  time.Time `json:"expire_time"`
	ServiceType int       `json:"service_type"`
}

//CheckAuthReq 登录验证
type CheckAuthReq struct {
	Code string `json:"code"` //code	string	是	跳转码(小商店服务市场跳转到第三方url里面会带上code)
}

//CheckAuthRsp ..
type CheckAuthRsp struct {
	APPID     string `json:"appid"`
	ServiceID int    `json:"service_id"`
}
