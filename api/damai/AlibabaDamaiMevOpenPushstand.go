package damai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenPushstand 大麦换验平台-第三方对外开放-看台接口pushStand
// alibaba.damai.mev.open.pushstand
//
// pushStand
func AlibabaDamaiMevOpenPushstand(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushstandAPIRequest, resp *damai.AlibabaDamaiMevOpenPushstandAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
