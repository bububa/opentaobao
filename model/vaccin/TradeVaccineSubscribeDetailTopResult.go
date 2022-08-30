package vaccin

// TradeVaccineSubscribeDetailTopResult 结构体
type TradeVaccineSubscribeDetailTopResult struct {
	// 预约单模型
	DetailTopDtoList []TradeVaccineSubscribeDetailTopDto `json:"detail_top_dto_list,omitempty" xml:"detail_top_dto_list>trade_vaccine_subscribe_detail_top_dto,omitempty"`
	// 业务报错code
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 业务报错消息
	BizMessage string `json:"biz_message,omitempty" xml:"biz_message,omitempty"`
	// 业务成功状态
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}
