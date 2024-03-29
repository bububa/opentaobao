package promotion

import (
	"sync"
)

// CouponTemplateItemRequest 结构体
type CouponTemplateItemRequest struct {
	// 券圈品设置
	PromActSkuList []PromActSku `json:"prom_act_sku_list,omitempty" xml:"prom_act_sku_list>prom_act_sku,omitempty"`
	// 模板表主键id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// ump模板ID
	SourceId int64 `json:"source_id,omitempty" xml:"source_id,omitempty"`
	// 用户信息
	UserInfo *UserInfo `json:"user_info,omitempty" xml:"user_info,omitempty"`
}

var poolCouponTemplateItemRequest = sync.Pool{
	New: func() any {
		return new(CouponTemplateItemRequest)
	},
}

// GetCouponTemplateItemRequest() 从对象池中获取CouponTemplateItemRequest
func GetCouponTemplateItemRequest() *CouponTemplateItemRequest {
	return poolCouponTemplateItemRequest.Get().(*CouponTemplateItemRequest)
}

// ReleaseCouponTemplateItemRequest 释放CouponTemplateItemRequest
func ReleaseCouponTemplateItemRequest(v *CouponTemplateItemRequest) {
	v.PromActSkuList = v.PromActSkuList[:0]
	v.Id = 0
	v.SourceId = 0
	v.UserInfo = nil
	poolCouponTemplateItemRequest.Put(v)
}
