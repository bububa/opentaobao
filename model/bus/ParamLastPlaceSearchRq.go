package bus

import (
	"sync"
)

// ParamLastPlaceSearchRq 结构体
type ParamLastPlaceSearchRq struct {
	// 城市编码
	DepCityCode string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// 北京市
	DepCityName string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
}

var poolParamLastPlaceSearchRq = sync.Pool{
	New: func() any {
		return new(ParamLastPlaceSearchRq)
	},
}

// GetParamLastPlaceSearchRq() 从对象池中获取ParamLastPlaceSearchRq
func GetParamLastPlaceSearchRq() *ParamLastPlaceSearchRq {
	return poolParamLastPlaceSearchRq.Get().(*ParamLastPlaceSearchRq)
}

// ReleaseParamLastPlaceSearchRq 释放ParamLastPlaceSearchRq
func ReleaseParamLastPlaceSearchRq(v *ParamLastPlaceSearchRq) {
	v.DepCityCode = ""
	v.DepCityName = ""
	poolParamLastPlaceSearchRq.Put(v)
}
