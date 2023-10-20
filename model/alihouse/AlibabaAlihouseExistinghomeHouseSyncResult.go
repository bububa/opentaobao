package alihouse

// AlibabaalihouseexistinghomehousesyncResult 结构体
type AlibabaalihouseexistinghomehousesyncResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 操作结果
	Data *SyncHouseResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
