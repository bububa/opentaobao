package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenDeliveryorderBatchcreateAPIRequest 发货单创建批量接口 API请求
// taobao.qimen.deliveryorder.batchcreate
//
// taobao.qimen.deliveryorder.batchcreate
type TaobaoQimenDeliveryorderBatchcreateAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderBatchCreateRequest
}

// NewTaobaoQimenDeliveryorderBatchcreateRequest 初始化TaobaoQimenDeliveryorderBatchcreateAPIRequest对象
func NewTaobaoQimenDeliveryorderBatchcreateRequest() *TaobaoQimenDeliveryorderBatchcreateAPIRequest {
	return &TaobaoQimenDeliveryorderBatchcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderBatchcreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.deliveryorder.batchcreate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderBatchcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenDeliveryorderBatchcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenDeliveryorderBatchcreateAPIRequest) SetRequest(_request *DeliveryOrderBatchCreateRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenDeliveryorderBatchcreateAPIRequest) GetRequest() *DeliveryOrderBatchCreateRequest {
	return r._request
}
