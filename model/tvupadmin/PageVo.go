package tvupadmin

import (
	"sync"
)

// PageVo 结构体
type PageVo struct {
	// 内容列表
	RecordList []ChildNodeContentVo `json:"record_list,omitempty" xml:"record_list>child_node_content_vo,omitempty"`
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 单页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolPageVo = sync.Pool{
	New: func() any {
		return new(PageVo)
	},
}

// GetPageVo() 从对象池中获取PageVo
func GetPageVo() *PageVo {
	return poolPageVo.Get().(*PageVo)
}

// ReleasePageVo 释放PageVo
func ReleasePageVo(v *PageVo) {
	v.RecordList = v.RecordList[:0]
	v.Total = 0
	v.PageNo = 0
	v.PageSize = 0
	poolPageVo.Put(v)
}
