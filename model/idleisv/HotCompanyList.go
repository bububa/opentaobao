package idleisv

// HotCompanyList 结构体
type HotCompanyList struct {
	// 快递公司代码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 快递公司名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 快递公司id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
