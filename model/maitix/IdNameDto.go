package maitix

// IdNameDto 结构体
type IdNameDto struct {
	// 城市名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 城市id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
