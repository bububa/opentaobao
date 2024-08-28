package damai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenDeleteproject 大麦换验平台-第三方对外开放-项目接口deleteProject
// alibaba.damai.mev.open.deleteproject
//
// deleteProject
func AlibabaDamaiMevOpenDeleteproject(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeleteprojectAPIRequest, resp *damai.AlibabaDamaiMevOpenDeleteprojectAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
