package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsSignnameQueryAPIRequest 淘宝短信签名查询 API请求
// taobao.jst.sms.signname.query
//
// 淘宝短信签名查询
type TaobaoJstSmsSignnameQueryAPIRequest struct {
	model.Params
	// 签名查询的入参
	_querySmsSignRequest *TopQuerySmsSignRequest
}

// NewTaobaoJstSmsSignnameQueryRequest 初始化TaobaoJstSmsSignnameQueryAPIRequest对象
func NewTaobaoJstSmsSignnameQueryRequest() *TaobaoJstSmsSignnameQueryAPIRequest {
	return &TaobaoJstSmsSignnameQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsSignnameQueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.signname.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsSignnameQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQuerySmsSignRequest is QuerySmsSignRequest Setter
// 签名查询的入参
func (r *TaobaoJstSmsSignnameQueryAPIRequest) SetQuerySmsSignRequest(_querySmsSignRequest *TopQuerySmsSignRequest) error {
	r._querySmsSignRequest = _querySmsSignRequest
	r.Set("query_sms_sign_request", _querySmsSignRequest)
	return nil
}

// GetQuerySmsSignRequest QuerySmsSignRequest Getter
func (r TaobaoJstSmsSignnameQueryAPIRequest) GetQuerySmsSignRequest() *TopQuerySmsSignRequest {
	return r._querySmsSignRequest
}
