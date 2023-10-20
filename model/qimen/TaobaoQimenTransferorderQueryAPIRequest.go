package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimentransferorderqueryAPIRequest 调拨单查询 API请求
// taobao.qimen.transferorder.query
//
// 调拨单查询
type TaobaoqimentransferorderqueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoqimentransferorderqueryStruct
}

// NewTaobaoqimentransferorderqueryRequest 初始化TaobaoqimentransferorderqueryAPIRequest对象
func NewTaobaoqimentransferorderqueryRequest() *TaobaoqimentransferorderqueryAPIRequest {
	return &TaobaoqimentransferorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimentransferorderqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.transferorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimentransferorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimentransferorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimentransferorderqueryAPIRequest) SetRequest(_request *TaobaoqimentransferorderqueryStruct) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimentransferorderqueryAPIRequest) GetRequest() *TaobaoqimentransferorderqueryStruct {
	return r._request
}
