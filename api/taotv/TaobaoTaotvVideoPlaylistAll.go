package taotv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/taotv"
)

// TaobaoTaotvVideoPlaylistAll 获取播单列表
// taobao.taotv.video.playlist.all
//
// 根据牌照和视频源等获取播单列表
func TaobaoTaotvVideoPlaylistAll(ctx context.Context, clt *core.SDKClient, req *taotv.TaobaoTaotvVideoPlaylistAllAPIRequest, resp *taotv.TaobaoTaotvVideoPlaylistAllAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
