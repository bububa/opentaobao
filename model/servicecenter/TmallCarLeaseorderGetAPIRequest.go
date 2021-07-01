package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseorderGetAPIRequest
获取租赁订单信息 API请求
tmall.car.leaseorder.get

卖家通过供销平台获取代销商的订单信息，但是部分情况下网商银行订单号获取不到，需要提供接口或者工具给卖家 */
type TmallCarLeaseorderGetAPIRequest struct {
	model.Params
	// 订单号
	_orderId int64
}

// New
