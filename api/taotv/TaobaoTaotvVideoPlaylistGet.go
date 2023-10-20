package taotv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/taotv"
)

// TaobaoTaotvVideoPlaylistGet 根据频道ID获取频道下节目单以及当前播放
// taobao.taotv.video.playlist.get
//
// 根据频道ID获取频道下节目单以及当前播放
func TaobaoTaotvVideoPlaylistGet(clt *core.SDKClient, req *taotv.TaobaoTaotvVideoPlaylistGetAPIRequest, session string) (*taotv.TaobaoTaotvVideoPlaylistGetAPIResponse, error) {
	var resp taotv.TaobaoTaotvVideoPlaylistGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
