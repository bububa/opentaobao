package campus

import (
	"sync"
)

// TimePeriodQuery 结构体
type TimePeriodQuery struct {
	// 时间规则ID
	TimePeriodIdList []int64 `json:"time_period_id_list,omitempty" xml:"time_period_id_list>int64,omitempty"`
	// 序列号
	SnNo string `json:"sn_no,omitempty" xml:"sn_no,omitempty"`
}

var poolTimePeriodQuery = sync.Pool{
	New: func() any {
		return new(TimePeriodQuery)
	},
}

// GetTimePeriodQuery() 从对象池中获取TimePeriodQuery
func GetTimePeriodQuery() *TimePeriodQuery {
	return poolTimePeriodQuery.Get().(*TimePeriodQuery)
}

// ReleaseTimePeriodQuery 释放TimePeriodQuery
func ReleaseTimePeriodQuery(v *TimePeriodQuery) {
	v.TimePeriodIdList = v.TimePeriodIdList[:0]
	v.SnNo = ""
	poolTimePeriodQuery.Put(v)
}
