package btrip

// OpenIsvSearchRQ 结构体
type OpenIsvSearchRQ struct {
	// 第三方企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 商旅审批单id
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 业务类型，5：机票改签审批 6：机票退票审批
	BizCategory int64 `json:"biz_category,omitempty" xml:"biz_category,omitempty"`
}
