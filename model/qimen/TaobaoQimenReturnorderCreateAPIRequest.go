package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenreturnordercreateAPIRequest 退货入库单创建接口 API请求
// taobao.qimen.returnorder.create
//
// taobao.qimen.returnorder.create
type TaobaoqimenreturnordercreateAPIRequest struct {
	model.Params
	//
	_request *ReturnOrderCreateRequest
}

// NewTaobaoqimenreturnordercreateRequest 初始化TaobaoqimenreturnordercreateAPIRequest对象
func NewTaobaoqimenreturnordercreateRequest() *TaobaoqimenreturnordercreateAPIRequest {
	return &TaobaoqimenreturnordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenreturnordercreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.returnorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenreturnordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenreturnordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenreturnordercreateAPIRequest) SetRequest(_request *ReturnOrderCreateRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenreturnordercreateAPIRequest) GetRequest() *ReturnOrderCreateRequest {
	return r._request
}
