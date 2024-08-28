package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSalestarCreativesGet （新）批量获取创意
// taobao.simba.salestar.creatives.get
//
// 取得一个推广组的所有创意或者根据一个创意Id列表取得一组创意；&lt;br/&gt;如果同时提供了推广组Id和创意id列表，则优先使用推广组Id；
func TaobaoSimbaSalestarCreativesGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaSalestarCreativesGetAPIRequest, resp *simba.TaobaoSimbaSalestarCreativesGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
