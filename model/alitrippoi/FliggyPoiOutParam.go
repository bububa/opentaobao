package alitrippoi

import (
	"sync"
)

// FliggyPoiOutParam 结构体
type FliggyPoiOutParam struct {
	// 选择城市名
	CityNames []string `json:"city_names,omitempty" xml:"city_names>string,omitempty"`
	// 开始时间
	StartData string `json:"start_data,omitempty" xml:"start_data,omitempty"`
	// 城市码可不填
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 开始页数
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}

var poolFliggyPoiOutParam = sync.Pool{
	New: func() any {
		return new(FliggyPoiOutParam)
	},
}

// GetFliggyPoiOutParam() 从对象池中获取FliggyPoiOutParam
func GetFliggyPoiOutParam() *FliggyPoiOutParam {
	return poolFliggyPoiOutParam.Get().(*FliggyPoiOutParam)
}

// ReleaseFliggyPoiOutParam 释放FliggyPoiOutParam
func ReleaseFliggyPoiOutParam(v *FliggyPoiOutParam) {
	v.CityNames = v.CityNames[:0]
	v.StartData = ""
	v.CityCode = 0
	v.PageSize = 0
	v.PageNum = 0
	poolFliggyPoiOutParam.Put(v)
}
