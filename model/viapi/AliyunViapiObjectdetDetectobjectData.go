package viapi

// AliyunviapiobjectdetdetectobjectData 结构体
type AliyunviapiobjectdetdetectobjectData struct {
	// 人体检测框的集合
	Elements []Elements `json:"elements,omitempty" xml:"elements>elements,omitempty"`
	// 输入图片的高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 输入图片的宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
}
