package dmp

// Topic 结构体
type Topic struct {
	// 榜单下的模版对象数组
	Templates []Template `json:"templates,omitempty" xml:"templates>template,omitempty"`
	// 榜单名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 榜单描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 榜单id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
