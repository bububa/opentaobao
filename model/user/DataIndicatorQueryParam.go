package user

import (
	"sync"
)

// DataIndicatorQueryParam 结构体
type DataIndicatorQueryParam struct {
	// 指标
	Measure []string `json:"measure,omitempty" xml:"measure>string,omitempty"`
	// 过滤条件
	Filter string `json:"filter,omitempty" xml:"filter,omitempty"`
	// 日期类型
	DateType string `json:"date_type,omitempty" xml:"date_type,omitempty"`
	// 固定
	NameSpace string `json:"name_space,omitempty" xml:"name_space,omitempty"`
	// CGP绑定的数据银行品牌商ID
	SmartId int64 `json:"smart_id,omitempty" xml:"smart_id,omitempty"`
}

var poolDataIndicatorQueryParam = sync.Pool{
	New: func() any {
		return new(DataIndicatorQueryParam)
	},
}

// GetDataIndicatorQueryParam() 从对象池中获取DataIndicatorQueryParam
func GetDataIndicatorQueryParam() *DataIndicatorQueryParam {
	return poolDataIndicatorQueryParam.Get().(*DataIndicatorQueryParam)
}

// ReleaseDataIndicatorQueryParam 释放DataIndicatorQueryParam
func ReleaseDataIndicatorQueryParam(v *DataIndicatorQueryParam) {
	v.Measure = v.Measure[:0]
	v.Filter = ""
	v.DateType = ""
	v.NameSpace = ""
	v.SmartId = 0
	poolDataIndicatorQueryParam.Put(v)
}
