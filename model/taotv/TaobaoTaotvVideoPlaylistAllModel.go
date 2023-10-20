package taotv

import (
	"sync"
)

// TaobaoTaotvVideoPlaylistAllModel 结构体
type TaobaoTaotvVideoPlaylistAllModel struct {
	// playListName
	PlayListName string `json:"play_list_name,omitempty" xml:"play_list_name,omitempty"`
	// bgPic
	BgPic string `json:"bg_pic,omitempty" xml:"bg_pic,omitempty"`
	// modifyTime
	ModifyTime string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
	// createTime
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// playListId
	PlayListId int64 `json:"play_list_id,omitempty" xml:"play_list_id,omitempty"`
}

var poolTaobaoTaotvVideoPlaylistAllModel = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvVideoPlaylistAllModel)
	},
}

// GetTaobaoTaotvVideoPlaylistAllModel() 从对象池中获取TaobaoTaotvVideoPlaylistAllModel
func GetTaobaoTaotvVideoPlaylistAllModel() *TaobaoTaotvVideoPlaylistAllModel {
	return poolTaobaoTaotvVideoPlaylistAllModel.Get().(*TaobaoTaotvVideoPlaylistAllModel)
}

// ReleaseTaobaoTaotvVideoPlaylistAllModel 释放TaobaoTaotvVideoPlaylistAllModel
func ReleaseTaobaoTaotvVideoPlaylistAllModel(v *TaobaoTaotvVideoPlaylistAllModel) {
	v.PlayListName = ""
	v.BgPic = ""
	v.ModifyTime = ""
	v.CreateTime = ""
	v.PlayListId = 0
	poolTaobaoTaotvVideoPlaylistAllModel.Put(v)
}
