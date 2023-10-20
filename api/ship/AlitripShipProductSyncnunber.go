package ship

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ship"
)

// Alitripshipproductsyncnunber 船票班次变更回调
// alitrip.ship.product.syncnunber
//
// 船票班次变更回调
func Alitripshipproductsyncnunber(clt *core.SDKClient, req *ship.AlitripshipproductsyncnunberAPIRequest, session string) (*ship.AlitripshipproductsyncnunberAPIResponse, error) {
	var resp ship.AlitripshipproductsyncnunberAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
