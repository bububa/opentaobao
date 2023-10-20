package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenentryorderqueryAPIRequest 入库单查询接口 API请求
// taobao.qimen.entryorder.query
//
// ERP调用接口，查询入库单信息;
type TaobaoqimenentryorderqueryAPIRequest struct {
	model.Params
	//
	_request *EntryOrderQueryRequest
}

// NewTaobaoqimenentryorderqueryRequest 初始化TaobaoqimenentryorderqueryAPIRequest对象
func NewTaobaoqimenentryorderqueryRequest() *TaobaoqimenentryorderqueryAPIRequest {
	return &TaobaoqimenentryorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenentryorderqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.entryorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenentryorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenentryorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenentryorderqueryAPIRequest) SetRequest(_request *EntryOrderQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenentryorderqueryAPIRequest) GetRequest() *EntryOrderQueryRequest {
	return r._request
}
