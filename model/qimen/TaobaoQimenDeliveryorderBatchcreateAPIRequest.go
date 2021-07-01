package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenDeliveryorderBatchcreateAPIRequest
发货单创建批量接口 API请求
taobao.qimen.deliveryorder.batchcreate

ERP调用接口，将发货信息批量推送给WMS */
type TaobaoQimenDeliveryorderBatchcreateAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderBatchCreateRequest
}

// New
