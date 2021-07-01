package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenDeliveryorderQueryAPIRequest
发货单查询接口 API请求
taobao.qimen.deliveryorder.query

ERP调用奇门的发货单查询接口，查询发货单详情 */
type TaobaoQimenDeliveryorderQueryAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderQueryRequest
}

// New
