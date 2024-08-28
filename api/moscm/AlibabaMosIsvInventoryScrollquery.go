package moscm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// AlibabaMosIsvInventoryScrollquery 滚动查询库存数据
// alibaba.mos.isv.inventory.scrollquery
//
// 按专柜滚动查询有库存商品
func AlibabaMosIsvInventoryScrollquery(ctx context.Context, clt *core.SDKClient, req *moscm.AlibabaMosIsvInventoryScrollqueryAPIRequest, resp *moscm.AlibabaMosIsvInventoryScrollqueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
