package tbtrade

// LogisticsAgreement 结构体
type LogisticsAgreement struct {
	// 服务承诺/异常文案
	LogisticsServiceMsg string `json:"logistics_service_msg,omitempty" xml:"logistics_service_msg,omitempty"`
	// 物流服务业务身份
	AsdpBizType string `json:"asdp_biz_type,omitempty" xml:"asdp_biz_type,omitempty"`
	// 物流信息，多个值时用英文逗号隔开
	AsdpAds string `json:"asdp_ads,omitempty" xml:"asdp_ads,omitempty"`
	// 计划送达时间
	SignTime string `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
	// 承诺/最晚送达时间
	PromiseSignTime string `json:"promise_sign_time,omitempty" xml:"promise_sign_time,omitempty"`
	// ERP应推单时间(主单)
	PushTime string `json:"push_time,omitempty" xml:"push_time,omitempty"`
	// 1：不支持子单部分发货
	NoDetailPartConsign string `json:"no_detail_part_consign,omitempty" xml:"no_detail_part_consign,omitempty"`
	// 1：代表该订单是预售下沉订单类型为前置表达
	SinkType string `json:"sink_type,omitempty" xml:"sink_type,omitempty"`
}
