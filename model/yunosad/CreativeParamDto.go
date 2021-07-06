package yunosad

// CreativeParamDto 结构体
type CreativeParamDto struct {
	// 创意内容
	CreativeText []string `json:"creative_text,omitempty" xml:"creative_text>string,omitempty"`
	// 创意图片
	CreativeImageUrl []string `json:"creative_image_url,omitempty" xml:"creative_image_url>string,omitempty"`
	// 图标
	CreativeIconUrl []string `json:"creative_icon_url,omitempty" xml:"creative_icon_url>string,omitempty"`
	// 外部创意id
	CreativeId string `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
	// 外部创意名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 创意标题
	CreativeTitle string `json:"creative_title,omitempty" xml:"creative_title,omitempty"`
	// 落地页地址
	LandingUrl string `json:"landing_url,omitempty" xml:"landing_url,omitempty"`
	// 类型
	ActionType string `json:"action_type,omitempty" xml:"action_type,omitempty"`
	// 广告尺寸大小
	SizeCode string `json:"size_code,omitempty" xml:"size_code,omitempty"`
	// 创意模板id
	CreativeTemplateId int64 `json:"creative_template_id,omitempty" xml:"creative_template_id,omitempty"`
	// 创意类型
	CreativeType int64 `json:"creative_type,omitempty" xml:"creative_type,omitempty"`
}
