package ascp

// WarehouseCooperationQueryResponse 结构体
type WarehouseCooperationQueryResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回对象
	Data *WarehouseCooperationQueryResultDetail `json:"data,omitempty" xml:"data,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
