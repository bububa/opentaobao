package taotv

import (
	"sync"
)

// TaobaoTaotvVideoPlaylistPageData 结构体
type TaobaoTaotvVideoPlaylistPageData struct {
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 创建时间
	ModifyTime string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
	// 图片
	BgPic string `json:"bg_pic,omitempty" xml:"bg_pic,omitempty"`
	// 播单名称
	PlayListName string `json:"play_list_name,omitempty" xml:"play_list_name,omitempty"`
	// 播单id
	PlayListId int64 `json:"play_list_id,omitempty" xml:"play_list_id,omitempty"`
}

var poolTaobaoTaotvVideoPlaylistPageData = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvVideoPlaylistPageData)
	},
}

// GetTaobaoTaotvVideoPlaylistPageData() 从对象池中获取TaobaoTaotvVideoPlaylistPageData
func GetTaobaoTaotvVideoPlaylistPageData() *TaobaoTaotvVideoPlaylistPageData {
	return poolTaobaoTaotvVideoPlaylistPageData.Get().(*TaobaoTaotvVideoPlaylistPageData)
}

// ReleaseTaobaoTaotvVideoPlaylistPageData 释放TaobaoTaotvVideoPlaylistPageData
func ReleaseTaobaoTaotvVideoPlaylistPageData(v *TaobaoTaotvVideoPlaylistPageData) {
	v.CreateTime = ""
	v.ModifyTime = ""
	v.BgPic = ""
	v.PlayListName = ""
	v.PlayListId = 0
	poolTaobaoTaotvVideoPlaylistPageData.Put(v)
}
