package promotion

import (
	"sync"
)

// SingleBenefitRequest 结构体
type SingleBenefitRequest struct {
	// 事务id
	UniqueId string `json:"unique_id,omitempty" xml:"unique_id,omitempty"`
	// 调试线索
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 权益关联的活动ID
	BizId string `json:"biz_id,omitempty" xml:"biz_id,omitempty"`
	// 支持每日赢宝箱多个卡片类型，默认“游戏惊喜”卡片，可选“美妆”卡片beauty
	CardType string `json:"card_type,omitempty" xml:"card_type,omitempty"`
	// 权益类型    其中ALIPAY_COUPON 对应的type值是1
	BenefitType int64 `json:"benefit_type,omitempty" xml:"benefit_type,omitempty"`
	// 权益发放数量
	SendCount int64 `json:"send_count,omitempty" xml:"send_count,omitempty"`
	// 关联活动id
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 卖家用户id,　可通过接口taobao.user.seller.get获取
	PurchaserId int64 `json:"purchaser_id,omitempty" xml:"purchaser_id,omitempty"`
	// 活动详情id
	DetailId int64 `json:"detail_id,omitempty" xml:"detail_id,omitempty"`
}

var poolSingleBenefitRequest = sync.Pool{
	New: func() any {
		return new(SingleBenefitRequest)
	},
}

// GetSingleBenefitRequest() 从对象池中获取SingleBenefitRequest
func GetSingleBenefitRequest() *SingleBenefitRequest {
	return poolSingleBenefitRequest.Get().(*SingleBenefitRequest)
}

// ReleaseSingleBenefitRequest 释放SingleBenefitRequest
func ReleaseSingleBenefitRequest(v *SingleBenefitRequest) {
	v.UniqueId = ""
	v.TraceId = ""
	v.BizId = ""
	v.CardType = ""
	v.BenefitType = 0
	v.SendCount = 0
	v.RelationId = 0
	v.PurchaserId = 0
	v.DetailId = 0
	poolSingleBenefitRequest.Put(v)
}
