package alihouse

import (
	"sync"
)

// BusinessActivityDataDto 结构体
type BusinessActivityDataDto struct {
	// 分账数据
	AccountDivisionList []AccountDivisionRuleDto `json:"account_division_list,omitempty" xml:"account_division_list>account_division_rule_dto,omitempty"`
	// 外部活动ID
	OuterActivityId string `json:"outer_activity_id,omitempty" xml:"outer_activity_id,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 优惠时间
	DiscountValidTime string `json:"discount_valid_time,omitempty" xml:"discount_valid_time,omitempty"`
	// 楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 活动开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 小二联系方式
	Contact string `json:"contact,omitempty" xml:"contact,omitempty"`
	// 11111
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 是否开通POS 1-是 0-否
	IsOpenPos int64 `json:"is_open_pos,omitempty" xml:"is_open_pos,omitempty"`
	// 协议ID
	AgreementId int64 `json:"agreement_id,omitempty" xml:"agreement_id,omitempty"`
	// 主体ID
	MerchantOpenId int64 `json:"merchant_open_id,omitempty" xml:"merchant_open_id,omitempty"`
	// 结算类型
	SettlementType int64 `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	// 优惠模式
	DiscountMode int64 `json:"discount_mode,omitempty" xml:"discount_mode,omitempty"`
	// 优惠时间类型
	DiscountTimeType int64 `json:"discount_time_type,omitempty" xml:"discount_time_type,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 签章ID
	SignatureId int64 `json:"signature_id,omitempty" xml:"signature_id,omitempty"`
	// 结算账户ID
	SettleId int64 `json:"settle_id,omitempty" xml:"settle_id,omitempty"`
	// 是否确客
	IsConfirmGuest int64 `json:"is_confirm_guest,omitempty" xml:"is_confirm_guest,omitempty"`
}

var poolBusinessActivityDataDto = sync.Pool{
	New: func() any {
		return new(BusinessActivityDataDto)
	},
}

// GetBusinessActivityDataDto() 从对象池中获取BusinessActivityDataDto
func GetBusinessActivityDataDto() *BusinessActivityDataDto {
	return poolBusinessActivityDataDto.Get().(*BusinessActivityDataDto)
}

// ReleaseBusinessActivityDataDto 释放BusinessActivityDataDto
func ReleaseBusinessActivityDataDto(v *BusinessActivityDataDto) {
	v.AccountDivisionList = v.AccountDivisionList[:0]
	v.OuterActivityId = ""
	v.ActivityName = ""
	v.DiscountValidTime = ""
	v.OuterId = ""
	v.StartTime = ""
	v.EndTime = ""
	v.Contact = ""
	v.OuterStoreId = ""
	v.IsOpenPos = 0
	v.AgreementId = 0
	v.MerchantOpenId = 0
	v.SettlementType = 0
	v.DiscountMode = 0
	v.DiscountTimeType = 0
	v.ItemId = 0
	v.SignatureId = 0
	v.SettleId = 0
	v.IsConfirmGuest = 0
	poolBusinessActivityDataDto.Put(v)
}
