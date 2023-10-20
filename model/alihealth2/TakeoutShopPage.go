package alihealth2

import (
	"sync"
)

// TakeoutShopPage 结构体
type TakeoutShopPage struct {
	// 店铺信息列表
	TakeoutSummaryInfos []TakeoutShopSummaryInfo `json:"takeout_summary_infos,omitempty" xml:"takeout_summary_infos>takeout_shop_summary_info,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 当前页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页条数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 上一页页码
	PrevPage int64 `json:"prev_page,omitempty" xml:"prev_page,omitempty"`
	// 下一页页码
	NextPage int64 `json:"next_page,omitempty" xml:"next_page,omitempty"`
}

var poolTakeoutShopPage = sync.Pool{
	New: func() any {
		return new(TakeoutShopPage)
	},
}

// GetTakeoutShopPage() 从对象池中获取TakeoutShopPage
func GetTakeoutShopPage() *TakeoutShopPage {
	return poolTakeoutShopPage.Get().(*TakeoutShopPage)
}

// ReleaseTakeoutShopPage 释放TakeoutShopPage
func ReleaseTakeoutShopPage(v *TakeoutShopPage) {
	v.TakeoutSummaryInfos = v.TakeoutSummaryInfos[:0]
	v.TotalCount = 0
	v.Page = 0
	v.PageSize = 0
	v.TotalPage = 0
	v.PrevPage = 0
	v.NextPage = 0
	poolTakeoutShopPage.Put(v)
}
