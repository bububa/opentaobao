package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenorderpendingAPIRequest 单据挂起（恢复）接口 API请求
// taobao.qimen.order.pending
//
// ERP调用奇门的接口,挂起某些创建的单据;场景介绍：ERP主动发起挂起（恢复）某些创建的单据，如入库单、出库单、退货单等
type TaobaoqimenorderpendingAPIRequest struct {
	model.Params
	//
	_request *OrderPendingRequest
}

// NewTaobaoqimenorderpendingRequest 初始化TaobaoqimenorderpendingAPIRequest对象
func NewTaobaoqimenorderpendingRequest() *TaobaoqimenorderpendingAPIRequest {
	return &TaobaoqimenorderpendingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenorderpendingAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.order.pending"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenorderpendingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenorderpendingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenorderpendingAPIRequest) SetRequest(_request *OrderPendingRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenorderpendingAPIRequest) GetRequest() *OrderPendingRequest {
	return r._request
}
