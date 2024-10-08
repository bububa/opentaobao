package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpMaterialAccessallowed 物料准入判断
// taobao.universalbp.material.accessallowed
//
// 入参推广信息，出参相关主体是否可投放。如果投放了风控不准入的商品，无法正常投放。
func TaobaoUniversalbpMaterialAccessallowed(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpMaterialAccessallowedAPIRequest, resp *simba.TaobaoUniversalbpMaterialAccessallowedAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
