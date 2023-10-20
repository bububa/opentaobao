package ieagency

import (
	"sync"
)

// PassengerChangeFeeVo 结构体
type PassengerChangeFeeVo struct {
	// 改签升舱费(单位:分）
	ChangeUpgradeFee int64 `json:"change_upgrade_fee,omitempty" xml:"change_upgrade_fee,omitempty"`
	// 改签服务费(单位:分）
	ChangeServiceFee int64 `json:"change_service_fee,omitempty" xml:"change_service_fee,omitempty"`
	// 乘机人ID
	PassengerId int64 `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
}

var poolPassengerChangeFeeVo = sync.Pool{
	New: func() any {
		return new(PassengerChangeFeeVo)
	},
}

// GetPassengerChangeFeeVo() 从对象池中获取PassengerChangeFeeVo
func GetPassengerChangeFeeVo() *PassengerChangeFeeVo {
	return poolPassengerChangeFeeVo.Get().(*PassengerChangeFeeVo)
}

// ReleasePassengerChangeFeeVo 释放PassengerChangeFeeVo
func ReleasePassengerChangeFeeVo(v *PassengerChangeFeeVo) {
	v.ChangeUpgradeFee = 0
	v.ChangeServiceFee = 0
	v.PassengerId = 0
	poolPassengerChangeFeeVo.Put(v)
}
