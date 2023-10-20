package taotv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/taotv"
)

// TaobaoTaotvVideoPlaylistPage 分页获取所有播单
// taobao.taotv.video.playlist.page
//
// 获取所有播单信息（分页）
func TaobaoTaotvVideoPlaylistPage(clt *core.SDKClient, req *taotv.TaobaoTaotvVideoPlaylistPageAPIRequest, resp *taotv.TaobaoTaotvVideoPlaylistPageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
