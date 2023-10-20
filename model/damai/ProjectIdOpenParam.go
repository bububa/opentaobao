package damai

import (
	"sync"
)

// ProjectIdOpenParam 结构体
type ProjectIdOpenParam struct {
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 项目id
	ProjectId int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
}

var poolProjectIdOpenParam = sync.Pool{
	New: func() any {
		return new(ProjectIdOpenParam)
	},
}

// GetProjectIdOpenParam() 从对象池中获取ProjectIdOpenParam
func GetProjectIdOpenParam() *ProjectIdOpenParam {
	return poolProjectIdOpenParam.Get().(*ProjectIdOpenParam)
}

// ReleaseProjectIdOpenParam 释放ProjectIdOpenParam
func ReleaseProjectIdOpenParam(v *ProjectIdOpenParam) {
	v.SupplierSecret = ""
	v.ProjectId = 0
	v.SystemId = 0
	poolProjectIdOpenParam.Put(v)
}
