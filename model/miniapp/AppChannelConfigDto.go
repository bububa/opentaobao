package miniapp

// AppChannelConfigDto 结构体
type AppChannelConfigDto struct {
	// 配置url
	ConfigUrl string `json:"config_url,omitempty" xml:"config_url,omitempty"`
	// 扩展参数
	ExtProperties string `json:"ext_properties,omitempty" xml:"ext_properties,omitempty"`
	// 渠道id
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 卡片描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 卡片标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 服务商简称
	IsvSimpleName string `json:"isv_simple_name,omitempty" xml:"isv_simple_name,omitempty"`
	// 对应服务市场code
	ArticleCode string `json:"article_code,omitempty" xml:"article_code,omitempty"`
	// 小程序id
	MiniappId int64 `json:"miniapp_id,omitempty" xml:"miniapp_id,omitempty"`
	// 状态是否有效
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
