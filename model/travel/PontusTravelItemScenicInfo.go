package travel

// PontusTravelItemScenicInfo 结构体
type PontusTravelItemScenicInfo struct {
	// 结构化景点信息 景点结构化信息和文本描述二选一
	ScenicList []PontusTravelItemScenic `json:"scenic_list,omitempty" xml:"scenic_list>pontus_travel_item_scenic,omitempty"`
	// 门票套餐名称
	TicketPackageName string `json:"ticket_package_name,omitempty" xml:"ticket_package_name,omitempty"`
	// 景点描述文本，&lt;1500字符 景点结构化信息和文本描述二选一
	ScenicDesc string `json:"scenic_desc,omitempty" xml:"scenic_desc,omitempty"`
}
