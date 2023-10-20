package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayCreativeVideoUnbind 创意与视频解绑接口
// taobao.subway.creative.video.unbind
//
// 将创意与视频解绑
func TaobaoSubwayCreativeVideoUnbind(clt *core.SDKClient, req *simba.TaobaoSubwayCreativeVideoUnbindAPIRequest, resp *simba.TaobaoSubwayCreativeVideoUnbindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
