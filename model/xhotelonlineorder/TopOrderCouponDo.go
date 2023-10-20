package xhotelonlineorder

import (
	"sync"
)

// TopOrderCouponDo 结构体
type TopOrderCouponDo struct {
	// 卡劵名称
	CouponTitle string `json:"coupon_title,omitempty" xml:"coupon_title,omitempty"`
	// 卡劵订单号
	CouponOutBizId string `json:"coupon_out_biz_id,omitempty" xml:"coupon_out_biz_id,omitempty"`
	// 面额
	FaceValue string `json:"face_value,omitempty" xml:"face_value,omitempty"`
	// 单间夜面额
	RoomNightFaceValue string `json:"room_night_face_value,omitempty" xml:"room_night_face_value,omitempty"`
	// 卡劵说明
	Instruction string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// 卡劵模块id
	CouponBizType int64 `json:"coupon_biz_type,omitempty" xml:"coupon_biz_type,omitempty"`
	// 是否可拆分
	CanSplit bool `json:"can_split,omitempty" xml:"can_split,omitempty"`
}

var poolTopOrderCouponDo = sync.Pool{
	New: func() any {
		return new(TopOrderCouponDo)
	},
}

// GetTopOrderCouponDo() 从对象池中获取TopOrderCouponDo
func GetTopOrderCouponDo() *TopOrderCouponDo {
	return poolTopOrderCouponDo.Get().(*TopOrderCouponDo)
}

// ReleaseTopOrderCouponDo 释放TopOrderCouponDo
func ReleaseTopOrderCouponDo(v *TopOrderCouponDo) {
	v.CouponTitle = ""
	v.CouponOutBizId = ""
	v.FaceValue = ""
	v.RoomNightFaceValue = ""
	v.Instruction = ""
	v.CouponBizType = 0
	v.CanSplit = false
	poolTopOrderCouponDo.Put(v)
}
