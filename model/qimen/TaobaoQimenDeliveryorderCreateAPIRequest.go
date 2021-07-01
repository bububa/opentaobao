package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenDeliveryorderCreateAPIRequest
发货单创建接口 API请求
taobao.qimen.deliveryorder.create

taobao.qimen.deliveryorder.create */
type TaobaoQimenDeliveryorderCreateAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderCreateRequest
}

// New
