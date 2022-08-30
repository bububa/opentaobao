package usergrowth2

// MiniGameIndexDto 结构体
type MiniGameIndexDto struct {
	// 资产集合
	PropertyList []PropertyDto `json:"property_list,omitempty" xml:"property_list>property_dto,omitempty"`
}
