package tuanhotel

import (
	"sync"
)

// TopRootShopCategoryVoList 结构体
type TopRootShopCategoryVoList struct {
	// 二级类目列表
	ShopCategoryList []ShopCategoryVoList `json:"shop_category_list,omitempty" xml:"shop_category_list>shop_category_vo_list,omitempty"`
	// 一级类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 一级类目ID
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}

var poolTopRootShopCategoryVoList = sync.Pool{
	New: func() any {
		return new(TopRootShopCategoryVoList)
	},
}

// GetTopRootShopCategoryVoList() 从对象池中获取TopRootShopCategoryVoList
func GetTopRootShopCategoryVoList() *TopRootShopCategoryVoList {
	return poolTopRootShopCategoryVoList.Get().(*TopRootShopCategoryVoList)
}

// ReleaseTopRootShopCategoryVoList 释放TopRootShopCategoryVoList
func ReleaseTopRootShopCategoryVoList(v *TopRootShopCategoryVoList) {
	v.ShopCategoryList = v.ShopCategoryList[:0]
	v.CategoryName = ""
	v.CategoryId = 0
	poolTopRootShopCategoryVoList.Put(v)
}
