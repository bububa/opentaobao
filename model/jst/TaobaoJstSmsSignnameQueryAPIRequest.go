package jst

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstSmsSignnameQueryAPIRequest) Reset() {
	r._querySmsSignRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsSignnameQueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.signname.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsSignnameQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstSmsSignnameQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoJstSmsSignnameQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstSmsSignnameQueryRequest()
	},
}

// GetTaobaoJstSmsSignnameQueryRequest 从 sync.Pool 获取 TaobaoJstSmsSignnameQueryAPIRequest
func GetTaobaoJstSmsSignnameQueryAPIRequest() *TaobaoJstSmsSignnameQueryAPIRequest {
	return poolTaobaoJstSmsSignnameQueryAPIRequest.Get().(*TaobaoJstSmsSignnameQueryAPIRequest)
}

// ReleaseTaobaoJstSmsSignnameQueryAPIRequest 将 TaobaoJstSmsSignnameQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstSmsSignnameQueryAPIRequest(v *TaobaoJstSmsSignnameQueryAPIRequest) {
	v.Reset()
	poolTaobaoJstSmsSignnameQueryAPIRequest.Put(v)
}
