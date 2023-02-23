package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenReturnorderCreateAPIRequest 退货入库单创建接口 API请求
// taobao.qimen.returnorder.create
//
// ERP调用奇门的接口,创建退货单信息;该接口和入库单的区别就是该接口是从入库单接口中单独剥离出来的，专门处理退货引起的入 库操作
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
