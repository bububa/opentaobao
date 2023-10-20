package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// Alitriprentcarorderdetailquery 租车订单详情查询
// alitrip.rentcar.order.detail.query
//
// 租车订单详情查询
func Alitriprentcarorderdetailquery(clt *core.SDKClient, req *car.AlitriprentcarorderdetailqueryAPIRequest, session string) (*car.AlitriprentcarorderdetailqueryAPIResponse, error) {
	var resp car.AlitriprentcarorderdetailqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
