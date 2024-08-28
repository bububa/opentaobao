package damai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenPushperform 大麦换验平台-第三方对外开放-场次接口pushPerform
// alibaba.damai.mev.open.pushperform
//
// pushPerform
func AlibabaDamaiMevOpenPushperform(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushperformAPIRequest, resp *damai.AlibabaDamaiMevOpenPushperformAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
