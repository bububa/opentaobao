package alihouse

import (
	"sync"
)

// SubscriptionItemDto 结构体
type SubscriptionItemDto struct {
	// 外部支付方式id集合
	OuterPaymentIds []string `json:"outer_payment_ids,omitempty" xml:"outer_payment_ids>string,omitempty"`
	// 外部销售活动id
	OuterSalesActivityId string `json:"outer_sales_activity_id,omitempty" xml:"outer_sales_activity_id,omitempty"`
	// 外部楼盘id
	OuterProjectId string `json:"outer_project_id,omitempty" xml:"outer_project_id,omitempty"`
	// 外部项目店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 认购商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 认购商品楼盘页展示标题
	ShowTitle string `json:"show_title,omitempty" xml:"show_title,omitempty"`
	// 预存金商品名称
	PreDepositItemName string `json:"pre_deposit_item_name,omitempty" xml:"pre_deposit_item_name,omitempty"`
	// 预存金商品显示标题
	PreDepositItemTitle string `json:"pre_deposit_item_title,omitempty" xml:"pre_deposit_item_title,omitempty"`
	// 线上签约-网签时间值
	SignOnlineTimeVal string `json:"sign_online_time_val,omitempty" xml:"sign_online_time_val,omitempty"`
	// 结算账户关联的主体公司id
	ProjectCid int64 `json:"project_cid,omitempty" xml:"project_cid,omitempty"`
	// 认购商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 预存金商品id
	PreDepositItemId int64 `json:"pre_deposit_item_id,omitempty" xml:"pre_deposit_item_id,omitempty"`
	// 结算账户id
	SettleAid int64 `json:"settle_aid,omitempty" xml:"settle_aid,omitempty"`
	// 1- 线上  2-线下
	SignType int64 `json:"sign_type,omitempty" xml:"sign_type,omitempty"`
	// 签约超时时间(分钟)
	SignTimeoutTime int64 `json:"sign_timeout_time,omitempty" xml:"sign_timeout_time,omitempty"`
	// 线上签约-电子签章id
	ElecSignatureId int64 `json:"elec_signature_id,omitempty" xml:"elec_signature_id,omitempty"`
	// 线上签约-电子协议id
	ElecAgreementId int64 `json:"elec_agreement_id,omitempty" xml:"elec_agreement_id,omitempty"`
	// 线上签约-盖章方式
	SealType int64 `json:"seal_type,omitempty" xml:"seal_type,omitempty"`
	// 线上签约-网签时间方式
	SignOnlineTimeType int64 `json:"sign_online_time_type,omitempty" xml:"sign_online_time_type,omitempty"`
	// 限制购买数量
	BuyLimit int64 `json:"buy_limit,omitempty" xml:"buy_limit,omitempty"`
	// 是否异步销控
	IsAsyncSell int64 `json:"is_async_sell,omitempty" xml:"is_async_sell,omitempty"`
}

var poolSubscriptionItemDto = sync.Pool{
	New: func() any {
		return new(SubscriptionItemDto)
	},
}

// GetSubscriptionItemDto() 从对象池中获取SubscriptionItemDto
func GetSubscriptionItemDto() *SubscriptionItemDto {
	return poolSubscriptionItemDto.Get().(*SubscriptionItemDto)
}

// ReleaseSubscriptionItemDto 释放SubscriptionItemDto
func ReleaseSubscriptionItemDto(v *SubscriptionItemDto) {
	v.OuterPaymentIds = v.OuterPaymentIds[:0]
	v.OuterSalesActivityId = ""
	v.OuterProjectId = ""
	v.OuterStoreId = ""
	v.ItemName = ""
	v.ShowTitle = ""
	v.PreDepositItemName = ""
	v.PreDepositItemTitle = ""
	v.SignOnlineTimeVal = ""
	v.ProjectCid = 0
	v.ItemId = 0
	v.PreDepositItemId = 0
	v.SettleAid = 0
	v.SignType = 0
	v.SignTimeoutTime = 0
	v.ElecSignatureId = 0
	v.ElecAgreementId = 0
	v.SealType = 0
	v.SignOnlineTimeType = 0
	v.BuyLimit = 0
	v.IsAsyncSell = 0
	poolSubscriptionItemDto.Put(v)
}
