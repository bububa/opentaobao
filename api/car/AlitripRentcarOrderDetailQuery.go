package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// AlitripRentcarOrderDetailQuery 租车订单详情查询
// alitrip.rentcar.order.detail.query
//
// 租车订单详情查询
func AlitripRentcarOrderDetailQuery(clt *core.SDKClient, req *car.AlitripRentcarOrderDetailQueryAPIRequest, resp *car.AlitripRentcarOrderDetailQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
