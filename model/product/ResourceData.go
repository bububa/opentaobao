package product

import (
	"sync"
)

// ResourceData 结构体
type ResourceData struct {
	// 资源列表
	Datas []ResourceDataRecord `json:"datas,omitempty" xml:"datas>resource_data_record,omitempty"`
	// 资源名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 资源类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolResourceData = sync.Pool{
	New: func() any {
		return new(ResourceData)
	},
}

// GetResourceData() 从对象池中获取ResourceData
func GetResourceData() *ResourceData {
	return poolResourceData.Get().(*ResourceData)
}

// ReleaseResourceData 释放ResourceData
func ReleaseResourceData(v *ResourceData) {
	v.Datas = v.Datas[:0]
	v.Name = ""
	v.Type = 0
	poolResourceData.Put(v)
}
