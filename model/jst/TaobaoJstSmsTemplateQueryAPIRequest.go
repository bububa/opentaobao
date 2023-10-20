package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsTemplateQueryAPIRequest 淘宝短信模板查询 API请求
// taobao.jst.sms.template.query
//
// 淘宝短信模板查询
type TaobaoJstSmsTemplateQueryAPIRequest struct {
	model.Params
	// 查询短信模板的入参
	_querySmsTemplateRequest *TopQuerySmsTemplateRequest
}

// NewTaobaoJstSmsTemplateQueryRequest 初始化TaobaoJstSmsTemplateQueryAPIRequest对象
func NewTaobaoJstSmsTemplateQueryRequest() *TaobaoJstSmsTemplateQueryAPIRequest {
	return &TaobaoJstSmsTemplateQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstSmsTemplateQueryAPIRequest) Reset() {
	r._querySmsTemplateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsTemplateQueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.template.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsTemplateQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstSmsTemplateQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuerySmsTemplateRequest is QuerySmsTemplateRequest Setter
// 查询短信模板的入参
func (r *TaobaoJstSmsTemplateQueryAPIRequest) SetQuerySmsTemplateRequest(_querySmsTemplateRequest *TopQuerySmsTemplateRequest) error {
	r._querySmsTemplateRequest = _querySmsTemplateRequest
	r.Set("query_sms_template_request", _querySmsTemplateRequest)
	return nil
}

// GetQuerySmsTemplateRequest QuerySmsTemplateRequest Getter
func (r TaobaoJstSmsTemplateQueryAPIRequest) GetQuerySmsTemplateRequest() *TopQuerySmsTemplateRequest {
	return r._querySmsTemplateRequest
}

var poolTaobaoJstSmsTemplateQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstSmsTemplateQueryRequest()
	},
}

// GetTaobaoJstSmsTemplateQueryRequest 从 sync.Pool 获取 TaobaoJstSmsTemplateQueryAPIRequest
func GetTaobaoJstSmsTemplateQueryAPIRequest() *TaobaoJstSmsTemplateQueryAPIRequest {
	return poolTaobaoJstSmsTemplateQueryAPIRequest.Get().(*TaobaoJstSmsTemplateQueryAPIRequest)
}

// ReleaseTaobaoJstSmsTemplateQueryAPIRequest 将 TaobaoJstSmsTemplateQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstSmsTemplateQueryAPIRequest(v *TaobaoJstSmsTemplateQueryAPIRequest) {
	v.Reset()
	poolTaobaoJstSmsTemplateQueryAPIRequest.Put(v)
}
