package tvupadmin

import (
	"sync"
)

// ChildNodeContentVo 结构体
type ChildNodeContentVo struct {
	// 内容名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 图片地址
	Picurl string `json:"picurl,omitempty" xml:"picurl,omitempty"`
	// 内容业务类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 扩展字段
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 主键ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 所属类目ID
	NodeId int64 `json:"node_id,omitempty" xml:"node_id,omitempty"`
	// 排序
	Sort int64 `json:"sort,omitempty" xml:"sort,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolChildNodeContentVo = sync.Pool{
	New: func() any {
		return new(ChildNodeContentVo)
	},
}

// GetChildNodeContentVo() 从对象池中获取ChildNodeContentVo
func GetChildNodeContentVo() *ChildNodeContentVo {
	return poolChildNodeContentVo.Get().(*ChildNodeContentVo)
}

// ReleaseChildNodeContentVo 释放ChildNodeContentVo
func ReleaseChildNodeContentVo(v *ChildNodeContentVo) {
	v.Name = ""
	v.Picurl = ""
	v.Type = ""
	v.Extra = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.Id = 0
	v.NodeId = 0
	v.Sort = 0
	v.Status = 0
	poolChildNodeContentVo.Put(v)
}
