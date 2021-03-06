package homeai

// FeatureWallDto 结构体
type FeatureWallDto struct {
	// 方向四元数
	Fronts []string `json:"fronts,omitempty" xml:"fronts>string,omitempty"`
	// 坐标
	Pos []string `json:"pos,omitempty" xml:"pos>string,omitempty"`
	// roomid
	Room string `json:"room,omitempty" xml:"room,omitempty"`
	// 硬装类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 包围盒
	BoundingBox *BoundingBoxDto `json:"bounding_box,omitempty" xml:"bounding_box,omitempty"`
}
