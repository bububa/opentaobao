package flightuppc

// InsProductBaseParam 结构体
type InsProductBaseParam struct {
	// 干系人，用于保险团队联系
	StakeHolders []InsPersonParam `json:"stake_holders,omitempty" xml:"stake_holders>ins_person_param,omitempty"`
	// 被保人信息列表
	Insureds []InsPersonParam `json:"insureds,omitempty" xml:"insureds>ins_person_param,omitempty"`
	// 航段信息列表，示例：{&#34;endCity&#34;:&#34;成都&#34;, &#34;arrCityCode&#34;:&#34;510100&#34;, &#34;startTime&#34;:&#34;2022-10-26 13:10:00&#34;, &#34;arrAirport&#34;:&#34;TFU&#34;, &#34;startCity&#34;:&#34;珠海&#34;, &#34;depAirport&#34;:&#34;ZUH&#34;, &#34;flightNo&#34;:&#34;CA2678&#34;, &#34;endTime&#34;:&#34;2022-10-26 15:40:00&#34;, &#34;airlineName&#34;:&#34;国航&#34;, &#34;depCityCode&#34;:&#34;440400&#34;, &#34;ticketNo&#34;:&#34;123123424&#34;}
	AirTicketSegmentList []string `json:"air_ticket_segment_list,omitempty" xml:"air_ticket_segment_list>string,omitempty"`
	// 外部订单号，同一个用户购买的保险的outOrderId都是一样的
	OutOrderId int64 `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 保险产品id
	PremiumId int64 `json:"premium_id,omitempty" xml:"premium_id,omitempty"`
}
