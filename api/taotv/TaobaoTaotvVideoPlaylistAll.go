package taotv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/taotv"
)

// TaobaoTaotvVideoPlaylistAll 获取播单列表
// taobao.taotv.video.playlist.all
//
// 根据牌照和视频源等获取播单列表
func TaobaoTaotvVideoPlaylistAll(clt *core.SDKClient, req *taotv.TaobaoTaotvVideoPlaylistAllAPIRequest, session string) (*taotv.TaobaoTaotvVideoPlaylistAllAPIResponse, error) {
	var resp taotv.TaobaoTaotvVideoPlaylistAllAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
