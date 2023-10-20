package bus

import (
	"sync"
)

// CitySearchRp 结构体
type CitySearchRp struct {
	// 城市集合
	Citys []CityDto `json:"citys,omitempty" xml:"citys>city_dto,omitempty"`
	// 错误代码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCitySearchRp = sync.Pool{
	New: func() any {
		return new(CitySearchRp)
	},
}

// GetCitySearchRp() 从对象池中获取CitySearchRp
func GetCitySearchRp() *CitySearchRp {
	return poolCitySearchRp.Get().(*CitySearchRp)
}

// ReleaseCitySearchRp 释放CitySearchRp
func ReleaseCitySearchRp(v *CitySearchRp) {
	v.Citys = v.Citys[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolCitySearchRp.Put(v)
}
