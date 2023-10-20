package car

import (
	"sync"
)

// RentCarDepositInfo 结构体
type RentCarDepositInfo struct {
	// 免押节点
	DepositNodes []RentCarDepositNode `json:"deposit_nodes,omitempty" xml:"deposit_nodes>rent_car_deposit_node,omitempty"`
	// 车辆免押金额
	CarPreDeposit string `json:"car_pre_deposit,omitempty" xml:"car_pre_deposit,omitempty"`
	// 免押金额,
	DepositNum string `json:"deposit_num,omitempty" xml:"deposit_num,omitempty"`
	// 信用免押状态
	DepositStatus string `json:"deposit_status,omitempty" xml:"deposit_status,omitempty"`
	// 单押双押
	DepositType string `json:"deposit_type,omitempty" xml:"deposit_type,omitempty"`
	// 违章免押金额,
	LegalPreDeposit string `json:"legal_pre_deposit,omitempty" xml:"legal_pre_deposit,omitempty"`
	// 还车后{0}天自动释放免押额度
	CarDepositExpireTime int64 `json:"car_deposit_expire_time,omitempty" xml:"car_deposit_expire_time,omitempty"`
	// 国内租车违章押金解压超时时间
	LegalDepositExpireTime int64 `json:"legal_deposit_expire_time,omitempty" xml:"legal_deposit_expire_time,omitempty"`
	// 是否信用免押
	EnableDeposit bool `json:"enable_deposit,omitempty" xml:"enable_deposit,omitempty"`
}

var poolRentCarDepositInfo = sync.Pool{
	New: func() any {
		return new(RentCarDepositInfo)
	},
}

// GetRentCarDepositInfo() 从对象池中获取RentCarDepositInfo
func GetRentCarDepositInfo() *RentCarDepositInfo {
	return poolRentCarDepositInfo.Get().(*RentCarDepositInfo)
}

// ReleaseRentCarDepositInfo 释放RentCarDepositInfo
func ReleaseRentCarDepositInfo(v *RentCarDepositInfo) {
	v.DepositNodes = v.DepositNodes[:0]
	v.CarPreDeposit = ""
	v.DepositNum = ""
	v.DepositStatus = ""
	v.DepositType = ""
	v.LegalPreDeposit = ""
	v.CarDepositExpireTime = 0
	v.LegalDepositExpireTime = 0
	v.EnableDeposit = false
	poolRentCarDepositInfo.Put(v)
}
