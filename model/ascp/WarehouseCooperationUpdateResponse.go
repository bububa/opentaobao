package ascp

// WarehouseCooperationUpdateResponse 结构体
type WarehouseCooperationUpdateResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回内容
	Data *WarehouseCooperationResultDetail `json:"data,omitempty" xml:"data,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
