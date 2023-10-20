package btrip

import (
	"sync"
)

// SuggestRs 结构体
type SuggestRs struct {
	// 城市列表
	Cities []CityVo `json:"cities,omitempty" xml:"cities>city_vo,omitempty"`
	// 是否为邻近城市
	Nearby bool `json:"nearby,omitempty" xml:"nearby,omitempty"`
}

var poolSuggestRs = sync.Pool{
	New: func() any {
		return new(SuggestRs)
	},
}

// GetSuggestRs() 从对象池中获取SuggestRs
func GetSuggestRs() *SuggestRs {
	return poolSuggestRs.Get().(*SuggestRs)
}

// ReleaseSuggestRs 释放SuggestRs
func ReleaseSuggestRs(v *SuggestRs) {
	v.Cities = v.Cities[:0]
	v.Nearby = false
	poolSuggestRs.Put(v)
}
