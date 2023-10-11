package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenReturnorderCreateAPIRequest 退货入库单创建接口 API请求
// taobao.qimen.returnorder.create
//
// taobao.qimen.returnorder.create
type TaobaoQimenReturnorderCreateAPIRequest struct {
	model.Params
	//
	_request *ReturnOrderCreateRequest
}

// NewTaobaoQimenReturnorderCreateRequest 初始化TaobaoQimenReturnorderCreateAPIRequest对象
func NewTaobaoQimenReturnorderCreateRequest() *TaobaoQimenReturnorderCreateAPIRequest {
	return &TaobaoQimenReturnorderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenReturnorderCreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.returnorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenReturnorderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenReturnorderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenReturnorderCreateAPIRequest) SetRequest(_request *ReturnOrderCreateRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenReturnorderCreateAPIRequest) GetRequest() *ReturnOrderCreateRequest {
	return r._request
}
