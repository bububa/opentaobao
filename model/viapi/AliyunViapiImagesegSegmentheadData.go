package viapi

// AliyunViapiImagesegSegmentheadData 结构体
type AliyunViapiImagesegSegmentheadData struct {
	// 人体检测框的集合
	Elements []Elements `json:"elements,omitempty" xml:"elements>elements,omitempty"`
}
