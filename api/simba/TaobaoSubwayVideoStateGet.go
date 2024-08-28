package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayVideoStateGet 获取视频状态
// taobao.subway.video.state.get
//
// 获取已上传视频的状态
func TaobaoSubwayVideoStateGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSubwayVideoStateGetAPIRequest, resp *simba.TaobaoSubwayVideoStateGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
