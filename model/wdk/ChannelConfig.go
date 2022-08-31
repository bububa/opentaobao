package wdk

// ChannelConfig 结构体
type ChannelConfig struct {
	// 淘鲜达:&#34;31&#34;,饿了么:&#34;3&#34;,京东到家:&#34;26&#34;,美团外卖:&#34;2&#34;
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
}
