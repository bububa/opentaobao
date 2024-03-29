package fundplatform

import (
	"sync"
)

// CardFetchAsyncRequest 结构体
type CardFetchAsyncRequest struct {
	// 制卡详情
	CardFetchDetails []CardFetchDetailDto `json:"card_fetch_details,omitempty" xml:"card_fetch_details>card_fetch_detail_dto,omitempty"`
	// 购买主体类型 company 公司|person  个人
	BuyEntityType string `json:"buy_entity_type,omitempty" xml:"buy_entity_type,omitempty"`
	// 幂等字段，请保证唯一性,不要超过32位
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 售卖方式 online 线上|offline 线下
	SaleMode string `json:"sale_mode,omitempty" xml:"sale_mode,omitempty"`
	// 业务类型，由资金平台分配
	SubizType int64 `json:"subiz_type,omitempty" xml:"subiz_type,omitempty"`
	// 购买方在淘宝的用户ID
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 是否储值卡同步激活，默认为true
	Active bool `json:"active,omitempty" xml:"active,omitempty"`
}

var poolCardFetchAsyncRequest = sync.Pool{
	New: func() any {
		return new(CardFetchAsyncRequest)
	},
}

// GetCardFetchAsyncRequest() 从对象池中获取CardFetchAsyncRequest
func GetCardFetchAsyncRequest() *CardFetchAsyncRequest {
	return poolCardFetchAsyncRequest.Get().(*CardFetchAsyncRequest)
}

// ReleaseCardFetchAsyncRequest 释放CardFetchAsyncRequest
func ReleaseCardFetchAsyncRequest(v *CardFetchAsyncRequest) {
	v.CardFetchDetails = v.CardFetchDetails[:0]
	v.BuyEntityType = ""
	v.OutBizId = ""
	v.SaleMode = ""
	v.SubizType = 0
	v.UserId = 0
	v.Active = false
	poolCardFetchAsyncRequest.Put(v)
}
