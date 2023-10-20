package ship

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ship"
)

// AlitripShipProductSynccall 全量同步回调
// alitrip.ship.product.synccall
//
// 全量同步接口
func AlitripShipProductSynccall(clt *core.SDKClient, req *ship.AlitripShipProductSynccallAPIRequest, resp *ship.AlitripShipProductSynccallAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
