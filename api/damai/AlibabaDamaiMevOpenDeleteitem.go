package damai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenDeleteitem 大麦换验平台-第三方对外开放-票品接口deleteItem
// alibaba.damai.mev.open.deleteitem
//
// deleteItem
func AlibabaDamaiMevOpenDeleteitem(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeleteitemAPIRequest, resp *damai.AlibabaDamaiMevOpenDeleteitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
