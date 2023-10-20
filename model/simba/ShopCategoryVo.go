package simba

import (
	"sync"
)

// ShopCategoryVo 结构体
type ShopCategoryVo struct {
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 类目id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 父级类目id
	ParentCategoryId int64 `json:"parent_category_id,omitempty" xml:"parent_category_id,omitempty"`
}

var poolShopCategoryVo = sync.Pool{
	New: func() any {
		return new(ShopCategoryVo)
	},
}

// GetShopCategoryVo() 从对象池中获取ShopCategoryVo
func GetShopCategoryVo() *ShopCategoryVo {
	return poolShopCategoryVo.Get().(*ShopCategoryVo)
}

// ReleaseShopCategoryVo 释放ShopCategoryVo
func ReleaseShopCategoryVo(v *ShopCategoryVo) {
	v.CategoryName = ""
	v.CategoryId = 0
	v.ParentCategoryId = 0
	poolShopCategoryVo.Put(v)
}
