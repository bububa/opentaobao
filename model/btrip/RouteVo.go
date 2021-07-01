package btrip

// RouteVo 结构体
type RouteVo struct {
	// 来自城市
	FromCity string `json:"from_city,omitempty" xml:"from_city,omitempty"`
	// 到达城市
	ToCity string `json:"to_city,omitempty" xml:"to_city,omitempty"`
}
