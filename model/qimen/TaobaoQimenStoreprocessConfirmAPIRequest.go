package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenstoreprocessconfirmAPIRequest 仓内加工单确认接口 API请求
// taobao.qimen.storeprocess.confirm
//
// WMS调用奇门的接口,回传仓内加工单创建情况
type TaobaoqimenstoreprocessconfirmAPIRequest struct {
	model.Params
	//
	_request *StoreProcessConfirmRequest
}

// NewTaobaoqimenstoreprocessconfirmRequest 初始化TaobaoqimenstoreprocessconfirmAPIRequest对象
func NewTaobaoqimenstoreprocessconfirmRequest() *TaobaoqimenstoreprocessconfirmAPIRequest {
	return &TaobaoqimenstoreprocessconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenstoreprocessconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.storeprocess.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenstoreprocessconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenstoreprocessconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenstoreprocessconfirmAPIRequest) SetRequest(_request *StoreProcessConfirmRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenstoreprocessconfirmAPIRequest) GetRequest() *StoreProcessConfirmRequest {
	return r._request
}
