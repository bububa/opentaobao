package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenDeliveryorderBatchconfirmAPIRequest
发货单确认接口 API请求
taobao.qimen.deliveryorder.batchconfirm

taobao.qimen.deliveryorder.batchconfirm */
type TaobaoQimenDeliveryorderBatchconfirmAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderBatchConfirmRequest
}

// New
