package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimendeliveryorderbatchcreateAPIRequest 发货单创建批量接口 API请求
// taobao.qimen.deliveryorder.batchcreate
//
// taobao.qimen.deliveryorder.batchcreate
type TaobaoqimendeliveryorderbatchcreateAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderBatchCreateRequest
}

// NewTaobaoqimendeliveryorderbatchcreateRequest 初始化TaobaoqimendeliveryorderbatchcreateAPIRequest对象
func NewTaobaoqimendeliveryorderbatchcreateRequest() *TaobaoqimendeliveryorderbatchcreateAPIRequest {
	return &TaobaoqimendeliveryorderbatchcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimendeliveryorderbatchcreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.deliveryorder.batchcreate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimendeliveryorderbatchcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimendeliveryorderbatchcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimendeliveryorderbatchcreateAPIRequest) SetRequest(_request *DeliveryOrderBatchCreateRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimendeliveryorderbatchcreateAPIRequest) GetRequest() *DeliveryOrderBatchCreateRequest {
	return r._request
}
