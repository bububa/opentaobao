package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsTemplateCreateAPIRequest 淘宝短信模板创建 API请求
// taobao.jst.sms.template.create
//
// 聚石塔短信模板创建
type TaobaoJstSmsTemplateCreateAPIRequest struct {
	model.Params
	// 申请模板入参
	_smsTemplateForIsvRequest *AddSmsTemplateForIsvRequest
}

// NewTaobaoJstSmsTemplateCreateRequest 初始化TaobaoJstSmsTemplateCreateAPIRequest对象
func NewTaobaoJstSmsTemplateCreateRequest() *TaobaoJstSmsTemplateCreateAPIRequest {
	return &TaobaoJstSmsTemplateCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstSmsTemplateCreateAPIRequest) Reset() {
	r._smsTemplateForIsvRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsTemplateCreateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.template.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsTemplateCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstSmsTemplateCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSmsTemplateForIsvRequest is SmsTemplateForIsvRequest Setter
// 申请模板入参
func (r *TaobaoJstSmsTemplateCreateAPIRequest) SetSmsTemplateForIsvRequest(_smsTemplateForIsvRequest *AddSmsTemplateForIsvRequest) error {
	r._smsTemplateForIsvRequest = _smsTemplateForIsvRequest
	r.Set("sms_template_for_isv_request", _smsTemplateForIsvRequest)
	return nil
}

// GetSmsTemplateForIsvRequest SmsTemplateForIsvRequest Getter
func (r TaobaoJstSmsTemplateCreateAPIRequest) GetSmsTemplateForIsvRequest() *AddSmsTemplateForIsvRequest {
	return r._smsTemplateForIsvRequest
}

var poolTaobaoJstSmsTemplateCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstSmsTemplateCreateRequest()
	},
}

// GetTaobaoJstSmsTemplateCreateRequest 从 sync.Pool 获取 TaobaoJstSmsTemplateCreateAPIRequest
func GetTaobaoJstSmsTemplateCreateAPIRequest() *TaobaoJstSmsTemplateCreateAPIRequest {
	return poolTaobaoJstSmsTemplateCreateAPIRequest.Get().(*TaobaoJstSmsTemplateCreateAPIRequest)
}

// ReleaseTaobaoJstSmsTemplateCreateAPIRequest 将 TaobaoJstSmsTemplateCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstSmsTemplateCreateAPIRequest(v *TaobaoJstSmsTemplateCreateAPIRequest) {
	v.Reset()
	poolTaobaoJstSmsTemplateCreateAPIRequest.Put(v)
}
