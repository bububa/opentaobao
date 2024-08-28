package damai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenPushproject 大麦换验平台-第三方对外开放-项目接口pushProject
// alibaba.damai.mev.open.pushproject
//
// pushProject
func AlibabaDamaiMevOpenPushproject(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushprojectAPIRequest, resp *damai.AlibabaDamaiMevOpenPushprojectAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
