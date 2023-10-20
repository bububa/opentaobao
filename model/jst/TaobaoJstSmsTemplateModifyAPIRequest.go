package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsTemplateModifyAPIRequest 淘宝短信模板修改 API请求
// taobao.jst.sms.template.modify
//
// 淘宝短信模板修改
type TaobaoJstSmsTemplateModifyAPIRequest struct {
	model.Params
	// 修改模板的入参
	_modifySmsTemplateRequest *TopModifySmsTemplateRequest
}

// NewTaobaoJstSmsTemplateModifyRequest 初始化TaobaoJstSmsTemplateModifyAPIRequest对象
func NewTaobaoJstSmsTemplateModifyRequest() *TaobaoJstSmsTemplateModifyAPIRequest {
	return &TaobaoJstSmsTemplateModifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstSmsTemplateModifyAPIRequest) Reset() {
	r._modifySmsTemplateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsTemplateModifyAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.template.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsTemplateModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstSmsTemplateModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifySmsTemplateRequest is ModifySmsTemplateRequest Setter
// 修改模板的入参
func (r *TaobaoJstSmsTemplateModifyAPIRequest) SetModifySmsTemplateRequest(_modifySmsTemplateRequest *TopModifySmsTemplateRequest) error {
	r._modifySmsTemplateRequest = _modifySmsTemplateRequest
	r.Set("modify_sms_template_request", _modifySmsTemplateRequest)
	return nil
}

// GetModifySmsTemplateRequest ModifySmsTemplateRequest Getter
func (r TaobaoJstSmsTemplateModifyAPIRequest) GetModifySmsTemplateRequest() *TopModifySmsTemplateRequest {
	return r._modifySmsTemplateRequest
}

var poolTaobaoJstSmsTemplateModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstSmsTemplateModifyRequest()
	},
}

// GetTaobaoJstSmsTemplateModifyRequest 从 sync.Pool 获取 TaobaoJstSmsTemplateModifyAPIRequest
func GetTaobaoJstSmsTemplateModifyAPIRequest() *TaobaoJstSmsTemplateModifyAPIRequest {
	return poolTaobaoJstSmsTemplateModifyAPIRequest.Get().(*TaobaoJstSmsTemplateModifyAPIRequest)
}

// ReleaseTaobaoJstSmsTemplateModifyAPIRequest 将 TaobaoJstSmsTemplateModifyAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstSmsTemplateModifyAPIRequest(v *TaobaoJstSmsTemplateModifyAPIRequest) {
	v.Reset()
	poolTaobaoJstSmsTemplateModifyAPIRequest.Put(v)
}
