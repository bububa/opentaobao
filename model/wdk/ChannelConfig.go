package wdk

// ChannelConfig 结构体
type ChannelConfig struct {
	// 淘鲜达:"31",饿了么:"3",京东到家:"26",美团外卖:"2"
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
}
