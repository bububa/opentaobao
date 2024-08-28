package damai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenDeletevenue 大麦换验平台-第三方对外开放-场馆接口deleteVenue
// alibaba.damai.mev.open.deletevenue
//
// 开放接口，删除场馆
func AlibabaDamaiMevOpenDeletevenue(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeletevenueAPIRequest, resp *damai.AlibabaDamaiMevOpenDeletevenueAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
