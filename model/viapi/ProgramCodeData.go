package viapi

// ProgramCodeData 结构体
type ProgramCodeData struct {
	// 以图片左上角为坐标原点，小程序码区域左上角到y轴距离
	X int64 `json:"x,omitempty" xml:"x,omitempty"`
	// 以图片左上角为坐标原点，小程序码区域左上角到x轴距离
	Y int64 `json:"y,omitempty" xml:"y,omitempty"`
	// 小程序码区域宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 小程序码区域高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
}
