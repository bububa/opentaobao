package alsc

import (
	"sync"
)

// IssueVoucherReq 结构体
type IssueVoucherReq struct {
	// 活动id
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 品牌id(brandId和outerBrandId必传其一)
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 会员id
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 分享营销来源userId（需要用加密接口加密）
	FromUserId string `json:"from_user_id,omitempty" xml:"from_user_id,omitempty"`
	// 操作人id
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 操作人姓名
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 进店领券订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 外部品牌id
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// 门店id
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 幂等请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 进店领券订单来源
	OrderSrc int64 `json:"order_src,omitempty" xml:"order_src,omitempty"`
}

var poolIssueVoucherReq = sync.Pool{
	New: func() any {
		return new(IssueVoucherReq)
	},
}

// GetIssueVoucherReq() 从对象池中获取IssueVoucherReq
func GetIssueVoucherReq() *IssueVoucherReq {
	return poolIssueVoucherReq.Get().(*IssueVoucherReq)
}

// ReleaseIssueVoucherReq 释放IssueVoucherReq
func ReleaseIssueVoucherReq(v *IssueVoucherReq) {
	v.ActivityId = ""
	v.BrandId = ""
	v.CustomerId = ""
	v.FromUserId = ""
	v.OperatorId = ""
	v.OperatorName = ""
	v.OrderId = ""
	v.OutBrandId = ""
	v.OutShopId = ""
	v.RequestId = ""
	v.ShopId = ""
	v.OrderSrc = 0
	poolIssueVoucherReq.Put(v)
}
