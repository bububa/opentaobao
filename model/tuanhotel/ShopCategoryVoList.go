package tuanhotel

import (
	"sync"
)

// ShopCategoryVoList 结构体
type ShopCategoryVoList struct {
	// 二级类目名称
	CategoriesName string `json:"categories_name,omitempty" xml:"categories_name,omitempty"`
	// 二级类目ID
	CategoriesId int64 `json:"categories_id,omitempty" xml:"categories_id,omitempty"`
}

var poolShopCategoryVoList = sync.Pool{
	New: func() any {
		return new(ShopCategoryVoList)
	},
}

// GetShopCategoryVoList() 从对象池中获取ShopCategoryVoList
func GetShopCategoryVoList() *ShopCategoryVoList {
	return poolShopCategoryVoList.Get().(*ShopCategoryVoList)
}

// ReleaseShopCategoryVoList 释放ShopCategoryVoList
func ReleaseShopCategoryVoList(v *ShopCategoryVoList) {
	v.CategoriesName = ""
	v.CategoriesId = 0
	poolShopCategoryVoList.Put(v)
}
