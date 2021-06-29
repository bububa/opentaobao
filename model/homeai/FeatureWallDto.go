package homeai

// FeatureWallDTO 
type FeatureWallDTO struct {
    // 包围盒
    BoundingBox   *BoundingBoxDTO `json:"bounding_box,omitempty" xml:"bounding_box,omitempty"`
    // 方向四元数
    Fronts   []string `json:"fronts,omitempty" xml:"fronts>string,omitempty"`
    // 坐标
    Pos   []string `json:"pos,omitempty" xml:"pos>string,omitempty"`
    // roomid
    Room   string `json:"room,omitempty" xml:"room,omitempty"`
    // 硬装类型
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
}
