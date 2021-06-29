package fenxiao

// ProductImageList 
type ProductImageList struct {
    // 图片id
    ImageId   int64 `json:"image_id,omitempty" xml:"image_id,omitempty"`
    // 图片对应的url
    ImageUrl   string `json:"image_url,omitempty" xml:"image_url,omitempty"`
    // 图片顺序
    ImagePosition   int64 `json:"image_position,omitempty" xml:"image_position,omitempty"`
    // 图片类型（PRODUCT：产品图片  EXTPRODUCT：产品辅图  PROPERTIES：产品属性图片 ）
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    // 当图片类型为属性图片时，表示图片对应的属性pv对。
    Properties   string `json:"properties,omitempty" xml:"properties,omitempty"`
}
