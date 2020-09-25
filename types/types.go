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
