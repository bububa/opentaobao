package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// AlitripRentcarOrderDetailQuery 租车订单详情查询
// alitrip.rentcar.order.detail.query
//
// 租车订单详情查询
func AlitripRentcarOrderDetailQuery(clt *core.SDKClient, req *car.AlitripRentcarOrderDetailQueryAPIRequest, session string) (*car.AlitripRentcarOrderDetailQueryAPIResponse, error) {
	var resp car.AlitripRentcarOrderDetailQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
