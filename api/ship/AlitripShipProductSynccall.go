package ship

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ship"
)

// Alitripshipproductsynccall 全量同步回调
// alitrip.ship.product.synccall
//
// 全量同步接口
func Alitripshipproductsynccall(clt *core.SDKClient, req *ship.AlitripshipproductsynccallAPIRequest, session string) (*ship.AlitripshipproductsynccallAPIResponse, error) {
	var resp ship.AlitripshipproductsynccallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
