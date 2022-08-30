package maitix

// IdNameDto 结构体
type IdNameDto struct {
	// 国家名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 国家id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
