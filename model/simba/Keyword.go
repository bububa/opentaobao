package simba

// Keyword 结构体
type Keyword struct {
	// 主人昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 最后修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 关键词
	Word string `json:"word,omitempty" xml:"word,omitempty"`
	// 审核状态： &lt;br/&gt;audit_wait-待审核；&lt;br/&gt;audit_pass-审核通过(上线)；&lt;br/&gt;audit_reject-审核拒绝；&lt;br/&gt;audit_offline-审核直接下线；&lt;br/&gt;默认为 audit_pass。
	AuditStatus string `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// 审核拒绝原因描述
	AuditDesc string `json:"audit_desc,omitempty" xml:"audit_desc,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 匹配模式： 1 精确匹配，4 广泛匹配
	MatchScope string `json:"match_scope,omitempty" xml:"match_scope,omitempty"`
	// 词质量得分
	Qscore string `json:"qscore,omitempty" xml:"qscore,omitempty"`
	// 关键词id
	KeywordId int64 `json:"keyword_id,omitempty" xml:"keyword_id,omitempty"`
	// 推广组id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 推广计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 关键词出价，单位为分，不能小于5
	MaxPrice int64 `json:"max_price,omitempty" xml:"max_price,omitempty"`
	// 无线上是否采用PC*无线溢价的出价模式（1：是，0：否）
	MobileIsDefaultPrice int64 `json:"mobile_is_default_price,omitempty" xml:"mobile_is_default_price,omitempty"`
	// 无线独立出价（0：代表使用PC*无线溢价，如果不是0则代表无线出价的值）
	MaxMobilePrice int64 `json:"max_mobile_price,omitempty" xml:"max_mobile_price,omitempty"`
	// 是否使用推广组默认出价，true-是；false-否；
	IsDefaultPrice bool `json:"is_default_price,omitempty" xml:"is_default_price,omitempty"`
	// 是否是垃圾词，false-不是；true-是；垃圾词是近期无点击的词
	IsGarbage bool `json:"is_garbage,omitempty" xml:"is_garbage,omitempty"`
}
