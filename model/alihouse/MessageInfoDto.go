package alihouse

// MessageInfoDto 结构体
type MessageInfoDto struct {
	// 子账号
	BrokerId string `json:"broker_id,omitempty" xml:"broker_id,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 子业务类型
	BizSubType int64 `json:"biz_sub_type,omitempty" xml:"biz_sub_type,omitempty"`
	// 业务类型
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 顾问类型 1-置业顾问 2-经纪人
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
