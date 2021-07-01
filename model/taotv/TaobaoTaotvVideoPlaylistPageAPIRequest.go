package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTaotvVideoPlaylistPageAPIRequest
分页获取所有播单 API请求
taobao.taotv.video.playlist.page

获取所有播单信息（分页） */
type TaobaoTaotvVideoPlaylistPageAPIRequest struct {
	model.Params
	// 客户端信息
	_systemInfo string
	// 当前页（从1开始）
	_pageNo int64
}

// New
