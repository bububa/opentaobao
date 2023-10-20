package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwaycreativevideounbind 创意与视频解绑接口
// taobao.subway.creative.video.unbind
//
// 将创意与视频解绑
func Taobaosubwaycreativevideounbind(clt *core.SDKClient, req *simba.TaobaosubwaycreativevideounbindAPIRequest, session string) (*simba.TaobaosubwaycreativevideounbindAPIResponse, error) {
	var resp simba.TaobaosubwaycreativevideounbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
