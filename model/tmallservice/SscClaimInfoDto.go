package tmallservice

import (
	"sync"
)

// SscClaimInfoDto 结构体
type SscClaimInfoDto struct {
	// 保单号
	InsuranceOrderNo string `json:"insurance_order_no,omitempty" xml:"insurance_order_no,omitempty"`
	// 服务产品
	ServiceName string `json:"service_name,omitempty" xml:"service_name,omitempty"`
	// 商品名称
	AuctionName string `json:"auction_name,omitempty" xml:"auction_name,omitempty"`
	// 商品sku描述
	AuctionSkuDesc string `json:"auction_sku_desc,omitempty" xml:"auction_sku_desc,omitempty"`
	// 服务sku描述
	SkuDesc string `json:"sku_desc,omitempty" xml:"sku_desc,omitempty"`
	// 工单创建时间
	GmtCreateWorkcard string `json:"gmt_create_workcard,omitempty" xml:"gmt_create_workcard,omitempty"`
	// 工单完成时间
	CompletionDate string `json:"completion_date,omitempty" xml:"completion_date,omitempty"`
	// 问题描述、故障位置、故障判断
	ProblemDesc string `json:"problem_desc,omitempty" xml:"problem_desc,omitempty"`
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// sn
	Sn string `json:"sn,omitempty" xml:"sn,omitempty"`
	// 被保人
	Recognizee string `json:"recognizee,omitempty" xml:"recognizee,omitempty"`
	// 被保人社会统一号
	RecognizeeUnityNumber string `json:"recognizee_unity_number,omitempty" xml:"recognizee_unity_number,omitempty"`
	// 扩展属性 map的json格式
	TaskAttribute string `json:"task_attribute,omitempty" xml:"task_attribute,omitempty"`
	// 工单id
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
	// 淘宝交易实物子订单
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 服务订单编号
	ServiceOrderId int64 `json:"service_order_id,omitempty" xml:"service_order_id,omitempty"`
	// 商品id
	AuctionId int64 `json:"auction_id,omitempty" xml:"auction_id,omitempty"`
	// 服务单id
	SpServiceOrderId int64 `json:"sp_service_order_id,omitempty" xml:"sp_service_order_id,omitempty"`
	// 商品skuid
	AuctionSkuId int64 `json:"auction_sku_id,omitempty" xml:"auction_sku_id,omitempty"`
	// 理赔金额（分)
	ClaimFee int64 `json:"claim_fee,omitempty" xml:"claim_fee,omitempty"`
}

var poolSscClaimInfoDto = sync.Pool{
	New: func() any {
		return new(SscClaimInfoDto)
	},
}

// GetSscClaimInfoDto() 从对象池中获取SscClaimInfoDto
func GetSscClaimInfoDto() *SscClaimInfoDto {
	return poolSscClaimInfoDto.Get().(*SscClaimInfoDto)
}

// ReleaseSscClaimInfoDto 释放SscClaimInfoDto
func ReleaseSscClaimInfoDto(v *SscClaimInfoDto) {
	v.InsuranceOrderNo = ""
	v.ServiceName = ""
	v.AuctionName = ""
	v.AuctionSkuDesc = ""
	v.SkuDesc = ""
	v.GmtCreateWorkcard = ""
	v.CompletionDate = ""
	v.ProblemDesc = ""
	v.BuyerNick = ""
	v.Sn = ""
	v.Recognizee = ""
	v.RecognizeeUnityNumber = ""
	v.TaskAttribute = ""
	v.WorkcardId = 0
	v.BizOrderId = 0
	v.ServiceOrderId = 0
	v.AuctionId = 0
	v.SpServiceOrderId = 0
	v.AuctionSkuId = 0
	v.ClaimFee = 0
	poolSscClaimInfoDto.Put(v)
}
