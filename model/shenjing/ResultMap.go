package shenjing

// ResultMap 结构体
type ResultMap struct {
	// 图片URL
	PhotoUrl string `json:"photo_url,omitempty" xml:"photo_url,omitempty"`
	// 园区名称
	CampusName string `json:"campus_name,omitempty" xml:"campus_name,omitempty"`
}
