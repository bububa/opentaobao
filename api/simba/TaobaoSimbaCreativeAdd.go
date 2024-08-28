package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCreativeAdd 增加创意
// taobao.simba.creative.add
//
// 创建一个创意
func TaobaoSimbaCreativeAdd(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCreativeAddAPIRequest, resp *simba.TaobaoSimbaCreativeAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
