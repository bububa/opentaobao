package campus

import (
	"sync"
)

// CalenderTemplateQuery 结构体
type CalenderTemplateQuery struct {
	// 日历模版id集合
	TemplateIdList []int64 `json:"template_id_list,omitempty" xml:"template_id_list>int64,omitempty"`
	// 序列号
	SnNo string `json:"sn_no,omitempty" xml:"sn_no,omitempty"`
}

var poolCalenderTemplateQuery = sync.Pool{
	New: func() any {
		return new(CalenderTemplateQuery)
	},
}

// GetCalenderTemplateQuery() 从对象池中获取CalenderTemplateQuery
func GetCalenderTemplateQuery() *CalenderTemplateQuery {
	return poolCalenderTemplateQuery.Get().(*CalenderTemplateQuery)
}

// ReleaseCalenderTemplateQuery 释放CalenderTemplateQuery
func ReleaseCalenderTemplateQuery(v *CalenderTemplateQuery) {
	v.TemplateIdList = v.TemplateIdList[:0]
	v.SnNo = ""
	poolCalenderTemplateQuery.Put(v)
}
