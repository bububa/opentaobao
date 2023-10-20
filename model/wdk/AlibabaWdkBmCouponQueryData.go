package wdk

import (
	"sync"
)

// AlibabaWdkBmCouponQueryData 结构体
type AlibabaWdkBmCouponQueryData struct {
	// 补差商品列表
	PaymentItemDOList []PaymentItemDo `json:"payment_item_d_o_list,omitempty" xml:"payment_item_d_o_list>payment_item_do,omitempty"`
	// 券名称
	CouponName string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty"`
	// 发放开始时间
	ApplyStartTime string `json:"apply_start_time,omitempty" xml:"apply_start_time,omitempty"`
	// 发放结束时间
	ApplyEndTime string `json:"apply_end_time,omitempty" xml:"apply_end_time,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 补差比例
	PaymentRate string `json:"payment_rate,omitempty" xml:"payment_rate,omitempty"`
	// 券id
	CouponId int64 `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`
	// 满减门槛，单位为分
	StartFee int64 `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
	// 优惠金额，单位为分
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 补差类型，值为1表示全补至门店售价、值为2表示部分补固定金额、值为3表示补主档价、值为4表示自行约定、值为5表示按比例补差
	PaymentType int64 `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	// 补差渠道，值为1表示与淘鲜达结算、值为2表示与零售商结算
	PaymentChannel int64 `json:"payment_channel,omitempty" xml:"payment_channel,omitempty"`
}

var poolAlibabaWdkBmCouponQueryData = sync.Pool{
	New: func() any {
		return new(AlibabaWdkBmCouponQueryData)
	},
}

// GetAlibabaWdkBmCouponQueryData() 从对象池中获取AlibabaWdkBmCouponQueryData
func GetAlibabaWdkBmCouponQueryData() *AlibabaWdkBmCouponQueryData {
	return poolAlibabaWdkBmCouponQueryData.Get().(*AlibabaWdkBmCouponQueryData)
}

// ReleaseAlibabaWdkBmCouponQueryData 释放AlibabaWdkBmCouponQueryData
func ReleaseAlibabaWdkBmCouponQueryData(v *AlibabaWdkBmCouponQueryData) {
	v.PaymentItemDOList = v.PaymentItemDOList[:0]
	v.CouponName = ""
	v.ApplyStartTime = ""
	v.ApplyEndTime = ""
	v.GmtCreate = ""
	v.PaymentRate = ""
	v.CouponId = 0
	v.StartFee = 0
	v.Amount = 0
	v.PaymentType = 0
	v.PaymentChannel = 0
	poolAlibabaWdkBmCouponQueryData.Put(v)
}
