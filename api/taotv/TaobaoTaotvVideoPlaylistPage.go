package taotv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/taotv"
)

// TaobaoTaotvVideoPlaylistPage 分页获取所有播单
// taobao.taotv.video.playlist.page
//
// 获取所有播单信息（分页）
func TaobaoTaotvVideoPlaylistPage(ctx context.Context, clt *core.SDKClient, req *taotv.TaobaoTaotvVideoPlaylistPageAPIRequest, resp *taotv.TaobaoTaotvVideoPlaylistPageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
