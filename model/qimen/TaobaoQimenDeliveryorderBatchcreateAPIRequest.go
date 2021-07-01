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
func (r TaobaoQimenDeliveryorderBatchcreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
//
func (r *TaobaoQimenDeliveryorderBatchcreateAPIRequest) SetRequest(_request *DeliveryOrderBatchCreateRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenDeliveryorderBatchcreateAPIRequest) GetRequest() *DeliveryOrderBatchCreateRequest {
	return r._request
}
