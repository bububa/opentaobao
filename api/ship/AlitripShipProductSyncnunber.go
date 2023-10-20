package ship

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ship"
)

// AlitripShipProductSyncnunber 船票班次变更回调
// alitrip.ship.product.syncnunber
//
// 船票班次变更回调
func AlitripShipProductSyncnunber(clt *core.SDKClient, req *ship.AlitripShipProductSyncnunberAPIRequest, resp *ship.AlitripShipProductSyncnunberAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
