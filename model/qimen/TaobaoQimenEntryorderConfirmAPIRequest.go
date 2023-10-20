package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenentryorderconfirmAPIRequest 入库单确认接口 API请求
// taobao.qimen.entryorder.confirm
//
// WMS调用接口，回传入库单信息;
type TaobaoqimenentryorderconfirmAPIRequest struct {
	model.Params
	//
	_request *EntryOrderConfirmRequest
}

// NewTaobaoqimenentryorderconfirmRequest 初始化TaobaoqimenentryorderconfirmAPIRequest对象
func NewTaobaoqimenentryorderconfirmRequest() *TaobaoqimenentryorderconfirmAPIRequest {
	return &TaobaoqimenentryorderconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenentryorderconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.entryorder.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenentryorderconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenentryorderconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenentryorderconfirmAPIRequest) SetRequest(_request *EntryOrderConfirmRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenentryorderconfirmAPIRequest) GetRequest() *EntryOrderConfirmRequest {
	return r._request
}
