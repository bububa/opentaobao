package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// Taobaomediavideolist 获取商家视频列表
// taobao.media.video.list
//
// 用于获取授权商家的视频列表
func Taobaomediavideolist(clt *core.SDKClient, req *media.TaobaomediavideolistAPIRequest, session string) (*media.TaobaomediavideolistAPIResponse, error) {
	var resp media.TaobaomediavideolistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
