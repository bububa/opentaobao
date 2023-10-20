package idle

import (
	"sync"
)

// TenderItemListQry 结构体
type TenderItemListQry struct {
	// 商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 场次日期
	ScheduleDate string `json:"schedule_date,omitempty" xml:"schedule_date,omitempty"`
	// 场次编码（两位数）
	ScheduleNumber string `json:"schedule_number,omitempty" xml:"schedule_number,omitempty"`
	// 用户昵称
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 查询页码
	PageNumber int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 用户id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolTenderItemListQry = sync.Pool{
	New: func() any {
		return new(TenderItemListQry)
	},
}

// GetTenderItemListQry() 从对象池中获取TenderItemListQry
func GetTenderItemListQry() *TenderItemListQry {
	return poolTenderItemListQry.Get().(*TenderItemListQry)
}

// ReleaseTenderItemListQry 释放TenderItemListQry
func ReleaseTenderItemListQry(v *TenderItemListQry) {
	v.ItemId = ""
	v.ScheduleDate = ""
	v.ScheduleNumber = ""
	v.UserNick = ""
	v.PageNumber = 0
	v.PageSize = 0
	v.UserId = 0
	poolTenderItemListQry.Put(v)
}
