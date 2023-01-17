package drugtrace

// WesCodeRelationDto 结构体
type WesCodeRelationDto struct {
	// 存在上下级关系时返回下级码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 存在上下级关系时返回上级码
	ParentCode string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
}
