package alitripmerchant

import (
	"sync"
)

// DerbySpecialField 结构体
type DerbySpecialField struct {
	// 积分倒计时
	PointsExpirationDate string `json:"points_expiration_date,omitempty" xml:"points_expiration_date,omitempty"`
	// 会员卡倒计时
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
}

var poolDerbySpecialField = sync.Pool{
	New: func() any {
		return new(DerbySpecialField)
	},
}

// GetDerbySpecialField() 从对象池中获取DerbySpecialField
func GetDerbySpecialField() *DerbySpecialField {
	return poolDerbySpecialField.Get().(*DerbySpecialField)
}

// ReleaseDerbySpecialField 释放DerbySpecialField
func ReleaseDerbySpecialField(v *DerbySpecialField) {
	v.PointsExpirationDate = ""
	v.EndDate = ""
	poolDerbySpecialField.Put(v)
}
