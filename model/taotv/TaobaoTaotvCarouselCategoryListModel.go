package taotv

import (
	"sync"
)

// TaobaoTaotvCarouselCategoryListModel 结构体
type TaobaoTaotvCarouselCategoryListModel struct {
	// 分类频道列表
	ChannelList []Channels `json:"channel_list,omitempty" xml:"channel_list>channels,omitempty"`
	// 分类图片
	Pic string `json:"pic,omitempty" xml:"pic,omitempty"`
	// 分类名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 分类排序
	Sort int64 `json:"sort,omitempty" xml:"sort,omitempty"`
	// 分类牌照方
	Bcp int64 `json:"bcp,omitempty" xml:"bcp,omitempty"`
	// 分类ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolTaobaoTaotvCarouselCategoryListModel = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvCarouselCategoryListModel)
	},
}

// GetTaobaoTaotvCarouselCategoryListModel() 从对象池中获取TaobaoTaotvCarouselCategoryListModel
func GetTaobaoTaotvCarouselCategoryListModel() *TaobaoTaotvCarouselCategoryListModel {
	return poolTaobaoTaotvCarouselCategoryListModel.Get().(*TaobaoTaotvCarouselCategoryListModel)
}

// ReleaseTaobaoTaotvCarouselCategoryListModel 释放TaobaoTaotvCarouselCategoryListModel
func ReleaseTaobaoTaotvCarouselCategoryListModel(v *TaobaoTaotvCarouselCategoryListModel) {
	v.ChannelList = v.ChannelList[:0]
	v.Pic = ""
	v.Name = ""
	v.Sort = 0
	v.Bcp = 0
	v.Id = 0
	poolTaobaoTaotvCarouselCategoryListModel.Put(v)
}
