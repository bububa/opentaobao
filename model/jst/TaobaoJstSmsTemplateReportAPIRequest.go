package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsTemplateReportAPIRequest 服务商存量短信模板上传 API请求
// taobao.jst.sms.template.report
//
// 用于上传目前已经在阿里通信申请到的且正在使用的模板信息，确保模板数据正确，否则会导致短信发送失败！！！
type TaobaoJstSmsTemplateReportAPIRequest struct {
	model.Params
	// 存量短信模板上报入参
	_smsTemplateRequest *SmsTemplateRequest
}

// NewTaobaoJstSmsTemplateReportRequest 初始化TaobaoJstSmsTemplateReportAPIRequest对象
func NewTaobaoJstSmsTemplateReportRequest() *TaobaoJstSmsTemplateReportAPIRequest {
	return &TaobaoJstSmsTemplateReportAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstSmsTemplateReportAPIRequest) Reset() {
	r._smsTemplateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsTemplateReportAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.template.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsTemplateReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstSmsTemplateReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSmsTemplateRequest is SmsTemplateRequest Setter
// 存量短信模板上报入参
func (r *TaobaoJstSmsTemplateReportAPIRequest) SetSmsTemplateRequest(_smsTemplateRequest *SmsTemplateRequest) error {
	r._smsTemplateRequest = _smsTemplateRequest
	r.Set("sms_template_request", _smsTemplateRequest)
	return nil
}

// GetSmsTemplateRequest SmsTemplateRequest Getter
func (r TaobaoJstSmsTemplateReportAPIRequest) GetSmsTemplateRequest() *SmsTemplateRequest {
	return r._smsTemplateRequest
}

var poolTaobaoJstSmsTemplateReportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstSmsTemplateReportRequest()
	},
}

// GetTaobaoJstSmsTemplateReportRequest 从 sync.Pool 获取 TaobaoJstSmsTemplateReportAPIRequest
func GetTaobaoJstSmsTemplateReportAPIRequest() *TaobaoJstSmsTemplateReportAPIRequest {
	return poolTaobaoJstSmsTemplateReportAPIRequest.Get().(*TaobaoJstSmsTemplateReportAPIRequest)
}

// ReleaseTaobaoJstSmsTemplateReportAPIRequest 将 TaobaoJstSmsTemplateReportAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstSmsTemplateReportAPIRequest(v *TaobaoJstSmsTemplateReportAPIRequest) {
	v.Reset()
	poolTaobaoJstSmsTemplateReportAPIRequest.Put(v)
}
