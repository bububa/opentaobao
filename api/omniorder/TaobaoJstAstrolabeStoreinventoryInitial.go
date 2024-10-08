package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoJstAstrolabeStoreinventoryInitial 后端商品库存初始化
// taobao.jst.astrolabe.storeinventory.initial
//
// 初始化电商仓或门店库存，该接口一次可以初始化多个门店(或电商仓)的多个商品的多种类型库存。此接口只能使用一次，后续所有的库存变动都需走增量库存同步接口。
func TaobaoJstAstrolabeStoreinventoryInitial(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoJstAstrolabeStoreinventoryInitialAPIRequest, resp *omniorder.TaobaoJstAstrolabeStoreinventoryInitialAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
