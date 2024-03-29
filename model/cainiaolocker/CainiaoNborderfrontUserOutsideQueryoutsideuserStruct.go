package cainiaolocker

import (
	"sync"
)

// CainiaoNborderfrontUserOutsideQueryoutsideuserStruct 结构体
type CainiaoNborderfrontUserOutsideQueryoutsideuserStruct struct {
	// cpUserId
	CpUserId string `json:"cp_user_id,omitempty" xml:"cp_user_id,omitempty"`
	// cpCode
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 站点名称
	WorkStationName string `json:"work_station_name,omitempty" xml:"work_station_name,omitempty"`
	// 站点编码
	WorkStationCode string `json:"work_station_code,omitempty" xml:"work_station_code,omitempty"`
	// 支付宝账号
	AlipayAccount string `json:"alipay_account,omitempty" xml:"alipay_account,omitempty"`
	// 户名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 手机
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 杭州
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 321000
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
}

var poolCainiaoNborderfrontUserOutsideQueryoutsideuserStruct = sync.Pool{
	New: func() any {
		return new(CainiaoNborderfrontUserOutsideQueryoutsideuserStruct)
	},
}

// GetCainiaoNborderfrontUserOutsideQueryoutsideuserStruct() 从对象池中获取CainiaoNborderfrontUserOutsideQueryoutsideuserStruct
func GetCainiaoNborderfrontUserOutsideQueryoutsideuserStruct() *CainiaoNborderfrontUserOutsideQueryoutsideuserStruct {
	return poolCainiaoNborderfrontUserOutsideQueryoutsideuserStruct.Get().(*CainiaoNborderfrontUserOutsideQueryoutsideuserStruct)
}

// ReleaseCainiaoNborderfrontUserOutsideQueryoutsideuserStruct 释放CainiaoNborderfrontUserOutsideQueryoutsideuserStruct
func ReleaseCainiaoNborderfrontUserOutsideQueryoutsideuserStruct(v *CainiaoNborderfrontUserOutsideQueryoutsideuserStruct) {
	v.CpUserId = ""
	v.CpCode = ""
	v.WorkStationName = ""
	v.WorkStationCode = ""
	v.AlipayAccount = ""
	v.Name = ""
	v.Mobile = ""
	v.CityName = ""
	v.CityCode = ""
	poolCainiaoNborderfrontUserOutsideQueryoutsideuserStruct.Put(v)
}
