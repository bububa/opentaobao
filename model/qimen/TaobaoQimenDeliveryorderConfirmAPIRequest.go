package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenDeliveryorderConfirmAPIRequest
发货单确认接口 API请求
taobao.qimen.deliveryorder.confirm

taobao.qimen.deliveryorder.confirm */
type TaobaoQimenDeliveryorderConfirmAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderConfirmRequest
}

// New
