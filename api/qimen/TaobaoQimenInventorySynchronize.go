package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenInventorySynchronize 库存状态同步接口
// taobao.qimen.inventory.synchronize
//
// ERP通过该接口同步指定商品的库存信息
func TaobaoQimenInventorySynchronize(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenInventorySynchronizeAPIRequest, resp *qimen.TaobaoQimenInventorySynchronizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
