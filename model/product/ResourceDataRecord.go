package product

import (
	"sync"
)

// ResourceDataRecord 结构体
type ResourceDataRecord struct {
	// 扩展属性，包括城市名称，外部编码等
	ValueMap string `json:"value_map,omitempty" xml:"value_map,omitempty"`
	// 映射id
	MappingId int64 `json:"mapping_id,omitempty" xml:"mapping_id,omitempty"`
	// 项目id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 次数
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 资源类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 资源id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolResourceDataRecord = sync.Pool{
	New: func() any {
		return new(ResourceDataRecord)
	},
}

// GetResourceDataRecord() 从对象池中获取ResourceDataRecord
func GetResourceDataRecord() *ResourceDataRecord {
	return poolResourceDataRecord.Get().(*ResourceDataRecord)
}

// ReleaseResourceDataRecord 释放ResourceDataRecord
func ReleaseResourceDataRecord(v *ResourceDataRecord) {
	v.ValueMap = ""
	v.MappingId = 0
	v.ProductId = 0
	v.Num = 0
	v.Type = 0
	v.Id = 0
	poolResourceDataRecord.Put(v)
}
