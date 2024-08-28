package damai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenDeleteperform 大麦换验平台-第三方对外开放-场次接口deletePerform
// alibaba.damai.mev.open.deleteperform
//
// deletePerform
func AlibabaDamaiMevOpenDeleteperform(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeleteperformAPIRequest, resp *damai.AlibabaDamaiMevOpenDeleteperformAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
