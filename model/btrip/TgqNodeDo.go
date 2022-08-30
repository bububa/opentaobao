package btrip

// TgqNodeDo 结构体
type TgqNodeDo struct {
	// 规则列表
	DetailList []TimeNodeDo `json:"detail_list,omitempty" xml:"detail_list>time_node_do,omitempty"`
	// 是否支持
	Able bool `json:"able,omitempty" xml:"able,omitempty"`
}
