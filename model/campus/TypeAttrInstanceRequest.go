package campus

// TypeAttrInstanceRequest 结构体
type TypeAttrInstanceRequest struct {
	// 值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 业务属性code
	AttrCode string `json:"attr_code,omitempty" xml:"attr_code,omitempty"`
	// 空间单元或空间分组id
	SpaceId int64 `json:"space_id,omitempty" xml:"space_id,omitempty"`
	// 空间单元：1，空间分组：2
	Category int64 `json:"category,omitempty" xml:"category,omitempty"`
}
