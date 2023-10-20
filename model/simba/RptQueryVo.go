package simba

import (
	"sync"
)

// RptQueryVo 结构体
type RptQueryVo struct {
	// 查询条件列表;场景1、实体列表中只需填充汇总报表数据，此时条件列表传一个即可；场景2、实体列表中需同时查询两天的数据做对比，此时条件列表内传两天各自的查询条件；
	ConditionList []RptQueryConditionVo `json:"condition_list,omitempty" xml:"condition_list>rpt_query_condition_vo,omitempty"`
	// 查询指标字段
	Fields string `json:"fields,omitempty" xml:"fields,omitempty"`
}

var poolRptQueryVo = sync.Pool{
	New: func() any {
		return new(RptQueryVo)
	},
}

// GetRptQueryVo() 从对象池中获取RptQueryVo
func GetRptQueryVo() *RptQueryVo {
	return poolRptQueryVo.Get().(*RptQueryVo)
}

// ReleaseRptQueryVo 释放RptQueryVo
func ReleaseRptQueryVo(v *RptQueryVo) {
	v.ConditionList = v.ConditionList[:0]
	v.Fields = ""
	poolRptQueryVo.Put(v)
}
