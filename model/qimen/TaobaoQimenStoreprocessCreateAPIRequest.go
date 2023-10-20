package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenstoreprocesscreateAPIRequest 仓内加工单创建接口 API请求
// taobao.qimen.storeprocess.create
//
// ERP调用奇门的接口,创建仓内加工单
type TaobaoqimenstoreprocesscreateAPIRequest struct {
	model.Params
	//
	_request *StoreProcessCreateRequest
}

// NewTaobaoqimenstoreprocesscreateRequest 初始化TaobaoqimenstoreprocesscreateAPIRequest对象
func NewTaobaoqimenstoreprocesscreateRequest() *TaobaoqimenstoreprocesscreateAPIRequest {
	return &TaobaoqimenstoreprocesscreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenstoreprocesscreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.storeprocess.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenstoreprocesscreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenstoreprocesscreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenstoreprocesscreateAPIRequest) SetRequest(_request *StoreProcessCreateRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenstoreprocesscreateAPIRequest) GetRequest() *StoreProcessCreateRequest {
	return r._request
}
