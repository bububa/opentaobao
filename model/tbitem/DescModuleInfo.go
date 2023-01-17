package tbitem

// DescModuleInfo 结构体
type DescModuleInfo struct {
	// 代表宝贝描述中规范化打标使用到的模块id列表，以逗号分隔。
	AnchorModuleIds string `json:"anchor_module_ids,omitempty" xml:"anchor_module_ids,omitempty"`
	// 类型代表规范化打标的类型：人工或自动化
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
