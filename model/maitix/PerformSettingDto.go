package maitix

import (
	"sync"
)

// PerformSettingDto 结构体
type PerformSettingDto struct {
	// 入场方式 1纸质票入场 2电子票入场
	IssueEnterModesList []int64 `json:"issue_enter_modes_list,omitempty" xml:"issue_enter_modes_list>int64,omitempty"`
	// 出票方式 1纸质票 2静态二维码电子票 3动态二维码电子票 4身份证电子票 5 短信码电子票
	IssueTicketModesList []int64 `json:"issue_ticket_modes_list,omitempty" xml:"issue_ticket_modes_list>int64,omitempty"`
	// 取票方式,1-无纸化,2-快递票,3-自助换票,4-门店自取,这个决定了下单页用户看到的取票方式,如果是多个,可以让用户自由选择,可以参考大麦网提交订单后的页面
	TakeTicketTypes []int64 `json:"take_ticket_types,omitempty" xml:"take_ticket_types>int64,omitempty"`
	// 选座类型0-立即购买 1-选座购买,当是有座项目并且是列表有1.就可以h5选座
	SeatSelectTypeList []int64 `json:"seat_select_type_list,omitempty" xml:"seat_select_type_list>int64,omitempty"`
	// 证件类型(&#34;身份证&#34;-&#34;id_card&#34;,&#34;护照&#34;-&#34;passport&#34;,&#34;港澳居民来往内地通行证&#34;-&#34;hk_macao_pass&#34;,&#34;台湾居民来往大陆通行证&#34;-&#34;taiwan_compatriot_card&#34;,&#34;士兵／军官&#34;-&#34;soldier_officer_card&#34;)
	CardType string `json:"card_type,omitempty" xml:"card_type,omitempty"`
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 一单一证 0：不是，1：是
	IsOneOrderOneCard int64 `json:"is_one_order_one_card,omitempty" xml:"is_one_order_one_card,omitempty"`
	// 一票一证 0：不是，1：是
	IsOneTicketOneCard int64 `json:"is_one_ticket_one_card,omitempty" xml:"is_one_ticket_one_card,omitempty"`
	// 是否实名制入场 0：不是，1：是
	IsRealNameEnter int64 `json:"is_real_name_enter,omitempty" xml:"is_real_name_enter,omitempty"`
	// 销售设置 0开票 1预售
	SaleType int64 `json:"sale_type,omitempty" xml:"sale_type,omitempty"`
}

var poolPerformSettingDto = sync.Pool{
	New: func() any {
		return new(PerformSettingDto)
	},
}

// GetPerformSettingDto() 从对象池中获取PerformSettingDto
func GetPerformSettingDto() *PerformSettingDto {
	return poolPerformSettingDto.Get().(*PerformSettingDto)
}

// ReleasePerformSettingDto 释放PerformSettingDto
func ReleasePerformSettingDto(v *PerformSettingDto) {
	v.IssueEnterModesList = v.IssueEnterModesList[:0]
	v.IssueTicketModesList = v.IssueTicketModesList[:0]
	v.TakeTicketTypes = v.TakeTicketTypes[:0]
	v.SeatSelectTypeList = v.SeatSelectTypeList[:0]
	v.CardType = ""
	v.PerformId = 0
	v.IsOneOrderOneCard = 0
	v.IsOneTicketOneCard = 0
	v.IsRealNameEnter = 0
	v.SaleType = 0
	poolPerformSettingDto.Put(v)
}
