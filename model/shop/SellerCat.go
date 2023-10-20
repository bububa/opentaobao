package shop

import (
	"sync"
)

// SellerCat 结构体
type SellerCat struct {
	// 创建时间。格式：yyyy-MM-dd HH:mm:ss
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 卖家自定义类目名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 链接图片地址
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 店铺类目类型：可选值：manual_type：手动分类，new_type：新品上价， tree_type：二三级类目树 ，property_type：属性叶子类目树， brand_type：品牌推广
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 修改时间。格式：yyyy-MM-dd HH:mm:ss
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 卖家自定义类目编号
	Cid int64 `json:"cid,omitempty" xml:"cid,omitempty"`
	// 父类目编号，值等于0：表示此类目为店铺下的一级类目，值不等于0：表示此类目有父类目
	ParentCid int64 `json:"parent_cid,omitempty" xml:"parent_cid,omitempty"`
	// 该类目在页面上的排序位置
	SortOrder int64 `json:"sort_order,omitempty" xml:"sort_order,omitempty"`
}

var poolSellerCat = sync.Pool{
	New: func() any {
		return new(SellerCat)
	},
}

// GetSellerCat() 从对象池中获取SellerCat
func GetSellerCat() *SellerCat {
	return poolSellerCat.Get().(*SellerCat)
}

// ReleaseSellerCat 释放SellerCat
func ReleaseSellerCat(v *SellerCat) {
	v.Created = ""
	v.Name = ""
	v.PicUrl = ""
	v.Type = ""
	v.Modified = ""
	v.Cid = 0
	v.ParentCid = 0
	v.SortOrder = 0
	poolSellerCat.Put(v)
}
