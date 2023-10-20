package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimensuppliersynchronizeAPIRequest 供应商同步接口 API请求
// taobao.qimen.supplier.synchronize
//
// 这个接口用来同步供应商信息
type TaobaoqimensuppliersynchronizeAPIRequest struct {
	model.Params
	//
	_request *TaobaoqimensuppliersynchronizeRequest
}

// NewTaobaoqimensuppliersynchronizeRequest 初始化TaobaoqimensuppliersynchronizeAPIRequest对象
func NewTaobaoqimensuppliersynchronizeRequest() *TaobaoqimensuppliersynchronizeAPIRequest {
	return &TaobaoqimensuppliersynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimensuppliersynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.supplier.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimensuppliersynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimensuppliersynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimensuppliersynchronizeAPIRequest) SetRequest(_request *TaobaoqimensuppliersynchronizeRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimensuppliersynchronizeAPIRequest) GetRequest() *TaobaoqimensuppliersynchronizeRequest {
	return r._request
}
