package ticket

import (
	"sync"
)

// Scenic 结构体
type Scenic struct {
	// 收费项目列表
	ProductList []Product `json:"product_list,omitempty" xml:"product_list>product,omitempty"`
	// 标准景点ID
	AliScenicId string `json:"ali_scenic_id,omitempty" xml:"ali_scenic_id,omitempty"`
	// 标准景点名称
	AliScenicName string `json:"ali_scenic_name,omitempty" xml:"ali_scenic_name,omitempty"`
	// 商家景点ID
	OutScenicId string `json:"out_scenic_id,omitempty" xml:"out_scenic_id,omitempty"`
	// 商家景点名称
	OutScenicName string `json:"out_scenic_name,omitempty" xml:"out_scenic_name,omitempty"`
}

var poolScenic = sync.Pool{
	New: func() any {
		return new(Scenic)
	},
}

// GetScenic() 从对象池中获取Scenic
func GetScenic() *Scenic {
	return poolScenic.Get().(*Scenic)
}

// ReleaseScenic 释放Scenic
func ReleaseScenic(v *Scenic) {
	v.ProductList = v.ProductList[:0]
	v.AliScenicId = ""
	v.AliScenicName = ""
	v.OutScenicId = ""
	v.OutScenicName = ""
	poolScenic.Put(v)
}
