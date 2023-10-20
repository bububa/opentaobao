package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsTemplateDeleteAPIRequest 淘宝短信模板删除 API请求
// taobao.jst.sms.template.delete
//
// 淘宝短信模板删除
type TaobaoJstSmsTemplateDeleteAPIRequest struct {
	model.Params
	// 删除模板的入参
	_deleteSmsTemplateRequest *TopDeleteSmsTemplateRequest
}

// NewTaobaoJstSmsTemplateDeleteRequest 初始化TaobaoJstSmsTemplateDeleteAPIRequest对象
func NewTaobaoJstSmsTemplateDeleteRequest() *TaobaoJstSmsTemplateDeleteAPIRequest {
	return &TaobaoJstSmsTemplateDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstSmsTemplateDeleteAPIRequest) Reset() {
	r._deleteSmsTemplateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsTemplateDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.template.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsTemplateDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstSmsTemplateDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleteSmsTemplateRequest is DeleteSmsTemplateRequest Setter
// 删除模板的入参
func (r *TaobaoJstSmsTemplateDeleteAPIRequest) SetDeleteSmsTemplateRequest(_deleteSmsTemplateRequest *TopDeleteSmsTemplateRequest) error {
	r._deleteSmsTemplateRequest = _deleteSmsTemplateRequest
	r.Set("delete_sms_template_request", _deleteSmsTemplateRequest)
	return nil
}

// GetDeleteSmsTemplateRequest DeleteSmsTemplateRequest Getter
func (r TaobaoJstSmsTemplateDeleteAPIRequest) GetDeleteSmsTemplateRequest() *TopDeleteSmsTemplateRequest {
	return r._deleteSmsTemplateRequest
}

var poolTaobaoJstSmsTemplateDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstSmsTemplateDeleteRequest()
	},
}

// GetTaobaoJstSmsTemplateDeleteRequest 从 sync.Pool 获取 TaobaoJstSmsTemplateDeleteAPIRequest
func GetTaobaoJstSmsTemplateDeleteAPIRequest() *TaobaoJstSmsTemplateDeleteAPIRequest {
	return poolTaobaoJstSmsTemplateDeleteAPIRequest.Get().(*TaobaoJstSmsTemplateDeleteAPIRequest)
}

// ReleaseTaobaoJstSmsTemplateDeleteAPIRequest 将 TaobaoJstSmsTemplateDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstSmsTemplateDeleteAPIRequest(v *TaobaoJstSmsTemplateDeleteAPIRequest) {
	v.Reset()
	poolTaobaoJstSmsTemplateDeleteAPIRequest.Put(v)
}
