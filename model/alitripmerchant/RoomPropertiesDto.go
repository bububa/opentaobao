package alitripmerchant

// RoomPropertiesDto 结构体
type RoomPropertiesDto struct {
	// 设施类型
	SubType string `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
	// 设施名称
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}
