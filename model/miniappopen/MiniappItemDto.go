package miniappopen

import (
	"sync"
)

// MiniappItemDto 结构体
type MiniappItemDto struct {
	// 主图
	MainPicUrl string `json:"main_pic_url,omitempty" xml:"main_pic_url,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 外部商品id
	OutItemId string `json:"out_item_id,omitempty" xml:"out_item_id,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 一口价，单位【分】
	ReservePrice int64 `json:"reserve_price,omitempty" xml:"reserve_price,omitempty"`
	// 0-正常，其他均为不正常
	AuctionStatus int64 `json:"auction_status,omitempty" xml:"auction_status,omitempty"`
}

var poolMiniappItemDto = sync.Pool{
	New: func() any {
		return new(MiniappItemDto)
	},
}

// GetMiniappItemDto() 从对象池中获取MiniappItemDto
func GetMiniappItemDto() *MiniappItemDto {
	return poolMiniappItemDto.Get().(*MiniappItemDto)
}

// ReleaseMiniappItemDto 释放MiniappItemDto
func ReleaseMiniappItemDto(v *MiniappItemDto) {
	v.MainPicUrl = ""
	v.Title = ""
	v.OutItemId = ""
	v.ItemId = 0
	v.ReservePrice = 0
	v.AuctionStatus = 0
	poolMiniappItemDto.Put(v)
}
