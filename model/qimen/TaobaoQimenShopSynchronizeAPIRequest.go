package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenshopsynchronizeAPIRequest 店铺同步接口 API请求
// taobao.qimen.shop.synchronize
//
// 店铺同步接口描述
type TaobaoqimenshopsynchronizeAPIRequest struct {
	model.Params
	// 请求
	_request *TaobaoqimenshopsynchronizeRequest
}

// NewTaobaoqimenshopsynchronizeRequest 初始化TaobaoqimenshopsynchronizeAPIRequest对象
func NewTaobaoqimenshopsynchronizeRequest() *TaobaoqimenshopsynchronizeAPIRequest {
	return &TaobaoqimenshopsynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenshopsynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.shop.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenshopsynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenshopsynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
// 请求
func (r *TaobaoqimenshopsynchronizeAPIRequest) SetRequest(_request *TaobaoqimenshopsynchronizeRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenshopsynchronizeAPIRequest) GetRequest() *TaobaoqimenshopsynchronizeRequest {
	return r._request
}
