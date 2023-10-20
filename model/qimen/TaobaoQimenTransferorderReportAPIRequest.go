package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimentransferorderreportAPIRequest 调拨单通知 API请求
// taobao.qimen.transferorder.report
//
// 调拨单通知
type TaobaoqimentransferorderreportAPIRequest struct {
	model.Params
	//
	_request *TaobaoqimentransferorderreportStruct
}

// NewTaobaoqimentransferorderreportRequest 初始化TaobaoqimentransferorderreportAPIRequest对象
func NewTaobaoqimentransferorderreportRequest() *TaobaoqimentransferorderreportAPIRequest {
	return &TaobaoqimentransferorderreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimentransferorderreportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.transferorder.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimentransferorderreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimentransferorderreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimentransferorderreportAPIRequest) SetRequest(_request *TaobaoqimentransferorderreportStruct) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimentransferorderreportAPIRequest) GetRequest() *TaobaoqimentransferorderreportStruct {
	return r._request
}
