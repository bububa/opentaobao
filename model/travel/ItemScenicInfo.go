package travel

// ItemScenicInfo 结构体
type ItemScenicInfo struct {
	// 结构化景点信息 景点结构化信息和文本描述二选一
	ScenicList []ItemScenic `json:"scenic_list,omitempty" xml:"scenic_list>item_scenic,omitempty"`
}
