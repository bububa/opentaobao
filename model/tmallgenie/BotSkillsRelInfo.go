package tmallgenie

// BotSkillsRelInfo 结构体
type BotSkillsRelInfo struct {
	// 结果集
	Results []BotSkillsRelInfo `json:"results,omitempty" xml:"results>bot_skills_rel_info,omitempty"`
	// 提供商集合
	ServiceProviders []ServiceProvider `json:"service_providers,omitempty" xml:"service_providers>service_provider,omitempty"`
	// 典型语句
	Samples string `json:"samples,omitempty" xml:"samples,omitempty"`
	// 长描述
	LongDesc string `json:"long_desc,omitempty" xml:"long_desc,omitempty"`
	// 技能图片地址
	IconImgUrl string `json:"icon_img_url,omitempty" xml:"icon_img_url,omitempty"`
	// 类别
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// 技能名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 唤醒词
	InvocationName string `json:"invocation_name,omitempty" xml:"invocation_name,omitempty"`
	// 技能图片地址
	IcoinImageUrl string `json:"icoin_image_url,omitempty" xml:"icoin_image_url,omitempty"`
	// 技能描述信息
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 典型例句(多个，以\t分隔)
	Sample string `json:"sample,omitempty" xml:"sample,omitempty"`
	// 总页数
	PageCount int64 `json:"page_count,omitempty" xml:"page_count,omitempty"`
	// 分页单位
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 应用id
	BotId int64 `json:"bot_id,omitempty" xml:"bot_id,omitempty"`
	// 技能Id
	SkillId int64 `json:"skill_id,omitempty" xml:"skill_id,omitempty"`
}
