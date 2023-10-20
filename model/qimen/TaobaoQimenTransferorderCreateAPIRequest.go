package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimentransferordercreateAPIRequest 调拨单创建 API请求
// taobao.qimen.transferorder.create
//
// 调拨单创建
type TaobaoqimentransferordercreateAPIRequest struct {
	model.Params
	//
	_request *TaobaoqimentransferordercreateStruct
}

// NewTaobaoqimentransferordercreateRequest 初始化TaobaoqimentransferordercreateAPIRequest对象
func NewTaobaoqimentransferordercreateRequest() *TaobaoqimentransferordercreateAPIRequest {
	return &TaobaoqimentransferordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimentransferordercreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.transferorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimentransferordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimentransferordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimentransferordercreateAPIRequest) SetRequest(_request *TaobaoqimentransferordercreateStruct) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimentransferordercreateAPIRequest) GetRequest() *TaobaoqimentransferordercreateStruct {
	return r._request
}
