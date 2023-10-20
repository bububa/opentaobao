package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimendeliveryorderqueryAPIRequest 发货单查询接口 API请求
// taobao.qimen.deliveryorder.query
//
// ERP调用奇门的发货单查询接口，查询发货单详情
type TaobaoqimendeliveryorderqueryAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderQueryRequest
}

// NewTaobaoqimendeliveryorderqueryRequest 初始化TaobaoqimendeliveryorderqueryAPIRequest对象
func NewTaobaoqimendeliveryorderqueryRequest() *TaobaoqimendeliveryorderqueryAPIRequest {
	return &TaobaoqimendeliveryorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimendeliveryorderqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.deliveryorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimendeliveryorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimendeliveryorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimendeliveryorderqueryAPIRequest) SetRequest(_request *DeliveryOrderQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimendeliveryorderqueryAPIRequest) GetRequest() *DeliveryOrderQueryRequest {
	return r._request
}
