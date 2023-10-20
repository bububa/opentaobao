package alihouse

// AlibabaalihouseexistinghomequeryhousebaseinfoResult 结构体
type AlibabaalihouseexistinghomequeryhousebaseinfoResult struct {
	// 200
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回实体类
	Data *SyncHouseBaseInfoDto `json:"data,omitempty" xml:"data,omitempty"`
	// true
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
