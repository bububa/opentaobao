package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenreturnorderconfirmAPIRequest 退货入库单确认接口 API请求
// taobao.qimen.returnorder.confirm
//
// taobao.qimen.returnorder.confirm
type TaobaoqimenreturnorderconfirmAPIRequest struct {
	model.Params
	//
	_request *ReturnOrderConfirmRequest
}

// NewTaobaoqimenreturnorderconfirmRequest 初始化TaobaoqimenreturnorderconfirmAPIRequest对象
func NewTaobaoqimenreturnorderconfirmRequest() *TaobaoqimenreturnorderconfirmAPIRequest {
	return &TaobaoqimenreturnorderconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenreturnorderconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.returnorder.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenreturnorderconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenreturnorderconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenreturnorderconfirmAPIRequest) SetRequest(_request *ReturnOrderConfirmRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenreturnorderconfirmAPIRequest) GetRequest() *ReturnOrderConfirmRequest {
	return r._request
}
