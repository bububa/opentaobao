package ship

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ship"
)

/* AlitripShipProductSyncbase
基础信息修改回调
alitrip.ship.product.syncbase

基础信息修改回调 */
func AlitripShipProductSyncbase(clt *core.SDKClient, req *ship.AlitripShipProductSyncbaseAPIRequest, session string) (*ship.AlitripShipProductSyncbaseAPIResponse, error) {
	var resp ship.AlitripShipProductSyncbaseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
