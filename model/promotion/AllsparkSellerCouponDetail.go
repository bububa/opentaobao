package promotion

import (
	"sync"
)

// AllsparkSellerCouponDetail 结构体
type AllsparkSellerCouponDetail struct {
	// 商品优惠券会有商品id集合
	ItemIds []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
	// 状态名称
	StatusName string `json:"status_name,omitempty" xml:"status_name,omitempty"`
	// 券失效时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// mtop 店铺链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 券生效时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 券对外ID
	SpreadId string `json:"spread_id,omitempty" xml:"spread_id,omitempty"`
	// 券名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 券类型
	CouponTypeName string `json:"coupon_type_name,omitempty" xml:"coupon_type_name,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 设置发券数量总数
	TotalCount string `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 店铺名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 券类型
	CouponType int64 `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`
	// 状态信息
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 面额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 卖家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 满足金额阀值  如订单满多少元才可用
	StartFee int64 `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
	// 模板限领数量
	PersonLimitCount int64 `json:"person_limit_count,omitempty" xml:"person_limit_count,omitempty"`
	// 保留数量
	ReserveCount int64 `json:"reserve_count,omitempty" xml:"reserve_count,omitempty"`
	// 已领取数量
	ApplyCount int64 `json:"apply_count,omitempty" xml:"apply_count,omitempty"`
}

var poolAllsparkSellerCouponDetail = sync.Pool{
	New: func() any {
		return new(AllsparkSellerCouponDetail)
	},
}

// GetAllsparkSellerCouponDetail() 从对象池中获取AllsparkSellerCouponDetail
func GetAllsparkSellerCouponDetail() *AllsparkSellerCouponDetail {
	return poolAllsparkSellerCouponDetail.Get().(*AllsparkSellerCouponDetail)
}

// ReleaseAllsparkSellerCouponDetail 释放AllsparkSellerCouponDetail
func ReleaseAllsparkSellerCouponDetail(v *AllsparkSellerCouponDetail) {
	v.ItemIds = v.ItemIds[:0]
	v.StatusName = ""
	v.EndTime = ""
	v.Url = ""
	v.StartTime = ""
	v.SpreadId = ""
	v.Title = ""
	v.CouponTypeName = ""
	v.SellerNick = ""
	v.TotalCount = ""
	v.ShopName = ""
	v.CouponType = 0
	v.Status = 0
	v.Amount = 0
	v.SellerId = 0
	v.StartFee = 0
	v.PersonLimitCount = 0
	v.ReserveCount = 0
	v.ApplyCount = 0
	poolAllsparkSellerCouponDetail.Put(v)
}
