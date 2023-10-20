package simba

// TaobaouniversalbpcreativepreaddTopResult 结构体
type TaobaouniversalbpcreativepreaddTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	PreAddItemCreativeVO *PreAddItemCreativeVo `json:"pre_add_item_creative_v_o,omitempty" xml:"pre_add_item_creative_v_o,omitempty"`
}
