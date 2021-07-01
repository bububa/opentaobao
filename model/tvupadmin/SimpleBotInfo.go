package tvupadmin

// SimpleBotInfo 结构体
type SimpleBotInfo struct {
	// 设备名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 设备id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
