package ship

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ship"
)

// Alitripshipproductsyncbase 基础信息修改回调
// alitrip.ship.product.syncbase
//
// 基础信息修改回调
func Alitripshipproductsyncbase(clt *core.SDKClient, req *ship.AlitripshipproductsyncbaseAPIRequest, session string) (*ship.AlitripshipproductsyncbaseAPIResponse, error) {
	var resp ship.AlitripshipproductsyncbaseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
