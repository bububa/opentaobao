package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSalestarCreativeAdd （新）新建创意
// taobao.simba.salestar.creative.add
//
// 创建一个创意
func TaobaoSimbaSalestarCreativeAdd(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaSalestarCreativeAddAPIRequest, resp *simba.TaobaoSimbaSalestarCreativeAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
