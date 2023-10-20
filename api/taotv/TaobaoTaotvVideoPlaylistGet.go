package taotv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/taotv"
)

// Taobaotaotvvideoplaylistget 根据频道ID获取频道下节目单以及当前播放
// taobao.taotv.video.playlist.get
//
// 根据频道ID获取频道下节目单以及当前播放
func Taobaotaotvvideoplaylistget(clt *core.SDKClient, req *taotv.TaobaotaotvvideoplaylistgetAPIRequest, session string) (*taotv.TaobaotaotvvideoplaylistgetAPIResponse, error) {
	var resp taotv.TaobaotaotvvideoplaylistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
