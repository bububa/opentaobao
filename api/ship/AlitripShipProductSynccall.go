package ship

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ship"
)

// AlitripShipProductSynccall 全量同步回调
// alitrip.ship.product.synccall
//
// 全量同步接口
func AlitripShipProductSynccall(clt *core.SDKClient, req *ship.AlitripShipProductSynccallAPIRequest, session string) (*ship.AlitripShipProductSynccallAPIResponse, error) {
	var resp ship.AlitripShipProductSynccallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
