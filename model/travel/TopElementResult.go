package travel

// TopElementResult 结构体
type TopElementResult struct {
	// 元素的外部商家编码
	ElementOuterId string `json:"element_outer_id,omitempty" xml:"element_outer_id,omitempty"`
	// 创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 修改时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 删除时间
	Deleted string `json:"deleted,omitempty" xml:"deleted,omitempty"`
}
