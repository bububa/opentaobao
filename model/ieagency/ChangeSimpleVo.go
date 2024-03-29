package ieagency

import (
	"sync"
)

// ChangeSimpleVo 结构体
type ChangeSimpleVo struct {
	// 总改签服务费(单位:分）
	TotalChangeServiceFee string `json:"total_change_service_fee,omitempty" xml:"total_change_service_fee,omitempty"`
	// 乘机人改签费用(单位:分)
	PassengerChangeFeeVos *PassengerChangeFeeVo `json:"passenger_change_fee_vos,omitempty" xml:"passenger_change_fee_vos,omitempty"`
	// 总改签升舱费(单位:分)
	TotalChangeUpgradeFee int64 `json:"total_change_upgrade_fee,omitempty" xml:"total_change_upgrade_fee,omitempty"`
	// 是否有成功改签过
	HadChanged bool `json:"had_changed,omitempty" xml:"had_changed,omitempty"`
}

var poolChangeSimpleVo = sync.Pool{
	New: func() any {
		return new(ChangeSimpleVo)
	},
}

// GetChangeSimpleVo() 从对象池中获取ChangeSimpleVo
func GetChangeSimpleVo() *ChangeSimpleVo {
	return poolChangeSimpleVo.Get().(*ChangeSimpleVo)
}

// ReleaseChangeSimpleVo 释放ChangeSimpleVo
func ReleaseChangeSimpleVo(v *ChangeSimpleVo) {
	v.TotalChangeServiceFee = ""
	v.PassengerChangeFeeVos = nil
	v.TotalChangeUpgradeFee = 0
	v.HadChanged = false
	poolChangeSimpleVo.Put(v)
}
