package taotv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/taotv"
)

// Taobaotaotvvideoplaylistall 获取播单列表
// taobao.taotv.video.playlist.all
//
// 根据牌照和视频源等获取播单列表
func Taobaotaotvvideoplaylistall(clt *core.SDKClient, req *taotv.TaobaotaotvvideoplaylistallAPIRequest, session string) (*taotv.TaobaotaotvvideoplaylistallAPIResponse, error) {
	var resp taotv.TaobaotaotvvideoplaylistallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
