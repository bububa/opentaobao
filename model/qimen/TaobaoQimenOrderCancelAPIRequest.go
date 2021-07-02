package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderCancelAPIRequest 单据取消接口 API请求
// taobao.qimen.order.cancel
//
// ERP调用奇门的接口,取消创建单据操作。场景介绍：ERP主动发起取消某些创建的单据。如入库单、出库单、退货单等；所有的场景
type TaobaoQimenOrderCancelAPIRequest struct {
	model.Params
	//
	_request *OrderCancelRequest
}

// NewTaobaoQimenOrderCancelRequest 初始化TaobaoQimenOrderCancelAPIRequest对象
func NewTaobaoQimenOrderCancelRequest() *TaobaoQimenOrderCancelAPIRequest {
	return &TaobaoQimenOrderCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderCancelAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
//
func (r *TaobaoQimenOrderCancelAPIRequest) SetRequest(_request *OrderCancelRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenOrderCancelAPIRequest) GetRequest() *OrderCancelRequest {
	return r._request
}
