package taotv

import (
	"sync"
)

// TaobaoTaotvVideoPlaylistPageModel 结构体
type TaobaoTaotvVideoPlaylistPageModel struct {
	// 播单对象
	DataList []TaobaoTaotvVideoPlaylistPageData `json:"data_list,omitempty" xml:"data_list>taobao_taotv_video_playlist_page_data,omitempty"`
	// 当前页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 此接口默认每次获取100条
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 播单总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 总共页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

var poolTaobaoTaotvVideoPlaylistPageModel = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvVideoPlaylistPageModel)
	},
}

// GetTaobaoTaotvVideoPlaylistPageModel() 从对象池中获取TaobaoTaotvVideoPlaylistPageModel
func GetTaobaoTaotvVideoPlaylistPageModel() *TaobaoTaotvVideoPlaylistPageModel {
	return poolTaobaoTaotvVideoPlaylistPageModel.Get().(*TaobaoTaotvVideoPlaylistPageModel)
}

// ReleaseTaobaoTaotvVideoPlaylistPageModel 释放TaobaoTaotvVideoPlaylistPageModel
func ReleaseTaobaoTaotvVideoPlaylistPageModel(v *TaobaoTaotvVideoPlaylistPageModel) {
	v.DataList = v.DataList[:0]
	v.PageNo = 0
	v.PageSize = 0
	v.TotalCount = 0
	v.TotalPage = 0
	poolTaobaoTaotvVideoPlaylistPageModel.Put(v)
}
