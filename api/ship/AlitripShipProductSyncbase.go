package ship

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ship"
)

// AlitripShipProductSyncbase 基础信息修改回调
// alitrip.ship.product.syncbase
//
// 基础信息修改回调
func AlitripShipProductSyncbase(ctx context.Context, clt *core.SDKClient, req *ship.AlitripShipProductSyncbaseAPIRequest, resp *ship.AlitripShipProductSyncbaseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
