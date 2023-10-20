package promotion

import (
	"sync"
)

// SellerCouponDetail 结构体
type SellerCouponDetail struct {
	// 商品优惠券会有商品id集合
	ItemIds []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
	// 券名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 卖家ID
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 状态名称
	StatusName string `json:"status_name,omitempty" xml:"status_name,omitempty"`
	// 券类型
	CouponTypeName string `json:"coupon_type_name,omitempty" xml:"coupon_type_name,omitempty"`
	// 券生效时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 券失效时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// mtop 店铺链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 卖家名称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 券对外ID
	SpreadId string `json:"spread_id,omitempty" xml:"spread_id,omitempty"`
	// 店铺名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 面额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 设置发券数量总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 满足金额阀值  如订单满多少元才可用
	StartFee int64 `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
	// 券类型
	CouponType int64 `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`
	// 状态信息
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolSellerCouponDetail = sync.Pool{
	New: func() any {
		return new(SellerCouponDetail)
	},
}

// GetSellerCouponDetail() 从对象池中获取SellerCouponDetail
func GetSellerCouponDetail() *SellerCouponDetail {
	return poolSellerCouponDetail.Get().(*SellerCouponDetail)
}

// ReleaseSellerCouponDetail 释放SellerCouponDetail
func ReleaseSellerCouponDetail(v *SellerCouponDetail) {
	v.ItemIds = v.ItemIds[:0]
	v.Title = ""
	v.SellerId = ""
	v.StatusName = ""
	v.CouponTypeName = ""
	v.StartTime = ""
	v.EndTime = ""
	v.Url = ""
	v.SellerNick = ""
	v.SpreadId = ""
	v.ShopName = ""
	v.Amount = 0
	v.TotalCount = 0
	v.StartFee = 0
	v.CouponType = 0
	v.Status = 0
	poolSellerCouponDetail.Put(v)
}
