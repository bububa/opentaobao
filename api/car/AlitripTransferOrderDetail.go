package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// AlitripTransferOrderDetail 接送订单详情接口
// alitrip.transfer.order.detail
//
// 接送订单详情接口
func AlitripTransferOrderDetail(clt *core.SDKClient, req *car.AlitripTransferOrderDetailAPIRequest, resp *car.AlitripTransferOrderDetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
