package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimendeliveryordercreateAPIRequest 发货单创建接口 API请求
// taobao.qimen.deliveryorder.create
//
// taobao.qimen.deliveryorder.create
type TaobaoqimendeliveryordercreateAPIRequest struct {
	model.Params
	// 奇门仓储字段
	_request *DeliveryOrderCreateRequest
}

// NewTaobaoqimendeliveryordercreateRequest 初始化TaobaoqimendeliveryordercreateAPIRequest对象
func NewTaobaoqimendeliveryordercreateRequest() *TaobaoqimendeliveryordercreateAPIRequest {
	return &TaobaoqimendeliveryordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimendeliveryordercreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.deliveryorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimendeliveryordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimendeliveryordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
// 奇门仓储字段
func (r *TaobaoqimendeliveryordercreateAPIRequest) SetRequest(_request *DeliveryOrderCreateRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimendeliveryordercreateAPIRequest) GetRequest() *DeliveryOrderCreateRequest {
	return r._request
}
