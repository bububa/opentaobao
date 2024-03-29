package tvupadmin

import (
	"sync"
)

// ChildRecItemVo 结构体
type ChildRecItemVo struct {
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 业务类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 扩展数据
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 图片链接
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 自增ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 所属类目
	NodeId int64 `json:"node_id,omitempty" xml:"node_id,omitempty"`
	// 布局ID
	LayoutId int64 `json:"layout_id,omitempty" xml:"layout_id,omitempty"`
	// 状态：1-上线，0-下线
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 排序
	Sort int64 `json:"sort,omitempty" xml:"sort,omitempty"`
	// 主题ID
	RuleId int64 `json:"rule_id,omitempty" xml:"rule_id,omitempty"`
}

var poolChildRecItemVo = sync.Pool{
	New: func() any {
		return new(ChildRecItemVo)
	},
}

// GetChildRecItemVo() 从对象池中获取ChildRecItemVo
func GetChildRecItemVo() *ChildRecItemVo {
	return poolChildRecItemVo.Get().(*ChildRecItemVo)
}

// ReleaseChildRecItemVo 释放ChildRecItemVo
func ReleaseChildRecItemVo(v *ChildRecItemVo) {
	v.Name = ""
	v.Type = ""
	v.Extra = ""
	v.PicUrl = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.Id = 0
	v.Version = 0
	v.NodeId = 0
	v.LayoutId = 0
	v.Status = 0
	v.Sort = 0
	v.RuleId = 0
	poolChildRecItemVo.Put(v)
}
