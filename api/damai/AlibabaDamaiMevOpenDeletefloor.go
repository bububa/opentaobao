package damai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenDeletefloor 大麦换验平台-第三方对外开放-楼层接口deleteFloor
// alibaba.damai.mev.open.deletefloor
//
// deleteFloor
func AlibabaDamaiMevOpenDeletefloor(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeletefloorAPIRequest, resp *damai.AlibabaDamaiMevOpenDeletefloorAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
