package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// TaobaoMediaVideoList 获取商家视频列表
// taobao.media.video.list
//
// 用于获取授权商家的视频列表
func TaobaoMediaVideoList(clt *core.SDKClient, req *media.TaobaoMediaVideoListAPIRequest, resp *media.TaobaoMediaVideoListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
