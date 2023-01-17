package vaccin

// TradeVaccineSubscribeDetailTopResult 结构体
type TradeVaccineSubscribeDetailTopResult struct {
	// 预约单模型
	DetailTopDtoList []TradeVaccineSubscribeDetailTopDto `json:"detail_top_dto_list,omitempty" xml:"detail_top_dto_list>trade_vaccine_subscribe_detail_top_dto,omitempty"`
}
