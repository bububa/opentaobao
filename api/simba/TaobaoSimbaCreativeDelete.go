package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCreativeDelete 删除创意
// taobao.simba.creative.delete
//
// 删除一个创意
func TaobaoSimbaCreativeDelete(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCreativeDeleteAPIRequest, resp *simba.TaobaoSimbaCreativeDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
