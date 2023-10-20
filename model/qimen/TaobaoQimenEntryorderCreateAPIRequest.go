package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenentryordercreateAPIRequest 入库单创建接口 API请求
// taobao.qimen.entryorder.create
//
// taobao.qimen.entryorder.create
type TaobaoqimenentryordercreateAPIRequest struct {
	model.Params
	//
	_request *EntryOrderCreateRequest
}

// NewTaobaoqimenentryordercreateRequest 初始化TaobaoqimenentryordercreateAPIRequest对象
func NewTaobaoqimenentryordercreateRequest() *TaobaoqimenentryordercreateAPIRequest {
	return &TaobaoqimenentryordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenentryordercreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.entryorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenentryordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenentryordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenentryordercreateAPIRequest) SetRequest(_request *EntryOrderCreateRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenentryordercreateAPIRequest) GetRequest() *EntryOrderCreateRequest {
	return r._request
}
