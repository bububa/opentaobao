package viapi

// AliyunviapiimagesegsegmentheadData 结构体
type AliyunviapiimagesegsegmentheadData struct {
	// 人体检测框的集合
	Elements []Elements `json:"elements,omitempty" xml:"elements>elements,omitempty"`
}
