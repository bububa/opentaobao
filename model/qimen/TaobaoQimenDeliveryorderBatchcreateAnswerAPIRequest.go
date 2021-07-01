package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest
发货单创建结果通知接口(批量) API请求
taobao.qimen.deliveryorder.batchcreate.answer

WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回） */
type TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderBatchCreateAnswerRequest
}

// New
