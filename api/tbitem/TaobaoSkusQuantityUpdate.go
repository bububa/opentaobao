package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoSkusQuantityUpdate SKU库存修改
// taobao.skus.quantity.update
//
// 提供按照全量/增量的方式批量修改SKU库存的功能
func TaobaoSkusQuantityUpdate(ctx context.Context, clt *core.SDKClient, req *tbitem.TaobaoSkusQuantityUpdateAPIRequest, resp *tbitem.TaobaoSkusQuantityUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
