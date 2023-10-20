package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwaycreativevideobind 绑定视频到创意上
// taobao.subway.creative.video.bind
//
// 将用户上传的视频绑定到指定的创意上
func Taobaosubwaycreativevideobind(clt *core.SDKClient, req *simba.TaobaosubwaycreativevideobindAPIRequest, session string) (*simba.TaobaosubwaycreativevideobindAPIResponse, error) {
	var resp simba.TaobaosubwaycreativevideobindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
