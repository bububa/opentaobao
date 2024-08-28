package ship

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ship"
)

// AlitripShipProductSynccall 全量同步回调
// alitrip.ship.product.synccall
//
// 全量同步接口
func AlitripShipProductSynccall(ctx context.Context, clt *core.SDKClient, req *ship.AlitripShipProductSynccallAPIRequest, resp *ship.AlitripShipProductSynccallAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
