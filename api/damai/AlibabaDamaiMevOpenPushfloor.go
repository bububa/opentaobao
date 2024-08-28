package damai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenPushfloor 大麦换验平台-第三方对外开放-楼层接口pushFloor
// alibaba.damai.mev.open.pushfloor
//
// pushFloor
func AlibabaDamaiMevOpenPushfloor(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushfloorAPIRequest, resp *damai.AlibabaDamaiMevOpenPushfloorAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
