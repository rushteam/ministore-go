package ministore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client ..
type Client struct {
	// URL         string
	// AccessToken string
}

// Exec ..
func (c *Client) exec(url string, req interface{}, rsp interface{}) error {
	codec, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json",
		bytes.NewReader(codec))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, rsp)
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

// CheckAuthReq.Exec(req,resp)

// Response ..
type Response struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Data    interface{}
}

// Exec ..
func (c *CheckAuthReq) Exec(componentAccessToken string) (*CheckAuthRsp, error) {
	url := fmt.Sprintf(serviceCheckAuthURI, componentAccessToken)
	client.R().
	return &CheckAuthRsp{}, nil
}

//ServiceListReq 获取用户购买的在有效期内的服务列表
type ServiceListReq struct {
	Code string `json:"code"` //code	string	是	跳转码(小商店服务市场跳转到第三方url里面会带上code)
}

//Service ..
type Service struct {
	ServiceID   string    `json:"service_id"`
	ServiceName string    `json:"service_name"`
	ExpireTime  time.Time `json:"expire_time"`
	ServiceType int       `json:"service_type"`
}

// ServiceTypeProduct ..
const (
	ServiceTypeProduct int = 2 + iota
	ServiceTypOrder
	ServiceTypLogistics
)

// Exec ..
func (c *ServiceListReq) Exec(accessToken string) ([]*Service, error) {
	url := fmt.Sprintf(serviceListURI, accessToken)
	fmt.Println(url)
	var services []*Service
	return services, nil
}

// ServiceOrderListReq 获取用户购买的服务列表
type ServiceOrderListReq struct {
	StartCreateTime time.Time `json:"start_create_time"`
	EndCreateTime   time.Time `json:"end_create_time"`
	Page            int       `json:"page"`
	PageSize        int       `json:"page_size"`
}

// ServiceOrder 用户购买服务
type ServiceOrder struct {
	ServiceOrderID int `json:"service_order_id"`

	ServiceID   int       `json:"service_id"`
	ServiceName string    `json:"service_name"`
	ExpireTime  time.Time `json:"expire_time"`
	ServiceType int       `json:"service_type"`

	SpecificationID string  `json:"specification_id"`
	TotalPrice      float64 `json:"total_price"`
}

//CategoryListReq ..
type CategoryListReq struct {
	ParentCatID int `json:"f_cat_id"`
}

// Exec ..
func (c *CategoryListReq) Exec(accessToken string) ([]*Category, error) {
	url := fmt.Sprintf(serviceCheckAuthURI, accessToken)
	fmt.Println(url)
	//cat_list
	var catList []*Category
	return catList, nil
}

//Category ..
type Category struct {
	CatID       int    `json:"cat_id"`
	ParentCatID int    `json:"f_cat_id"`
	Name        string `json:"name"`
}
