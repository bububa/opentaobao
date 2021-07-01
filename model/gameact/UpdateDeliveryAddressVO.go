package gameact

// UpdateDeliveryAddressVO 结构体
type UpdateDeliveryAddressVO struct {
	// 是否成功更新或确认地址
	UpdateAddress bool `json:"update_address,omitempty" xml:"update_address,omitempty"`
}
