package tblogistics

// CheckDeliveryAuthTopResponse 结构体
type CheckDeliveryAuthTopResponse struct {
	// 扩展
	Features string `json:"features,omitempty" xml:"features,omitempty"`
	// 0正常,1异常
	AuthStatus int64 `json:"auth_status,omitempty" xml:"auth_status,omitempty"`
}
