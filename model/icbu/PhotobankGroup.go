package icbu

// PhotobankGroup 结构体
type PhotobankGroup struct {
	// 分组名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// level3
	Level3 int64 `json:"level3,omitempty" xml:"level3,omitempty"`
	// level2
	Level2 int64 `json:"level2,omitempty" xml:"level2,omitempty"`
	// level1
	Level1 int64 `json:"level1,omitempty" xml:"level1,omitempty"`
	// 分组id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
