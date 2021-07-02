package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderPendingAPIRequest 单据挂起（恢复）接口 API请求
// taobao.qimen.order.pending
//
// ERP调用奇门的接口,挂起某些创建的单据;场景介绍：ERP主动发起挂起（恢复）某些创建的单据，如入库单、出库单、退货单等
type TaobaoQimenOrderPendingAPIRequest struct {
	model.Params
	//
	_request *OrderPendingRequest
}

// NewTaobaoQimenOrderPendingRequest 初始化TaobaoQimenOrderPendingAPIRequest对象
func NewTaobaoQimenOrderPendingRequest() *TaobaoQimenOrderPendingAPIRequest {
	return &TaobaoQimenOrderPendingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderPendingAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.order.pending"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderPendingAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
//
func (r *TaobaoQimenOrderPendingAPIRequest) SetRequest(_request *OrderPendingRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenOrderPendingAPIRequest) GetRequest() *OrderPendingRequest {
	return r._request
}
