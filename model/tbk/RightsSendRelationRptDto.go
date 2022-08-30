package tbk

// RightsSendRelationRptDto 结构体
type RightsSendRelationRptDto struct {
	// 日期
	BizDate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 渠道关系id
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 红包发放数量
	FundNum int64 `json:"fund_num,omitempty" xml:"fund_num,omitempty"`
}
