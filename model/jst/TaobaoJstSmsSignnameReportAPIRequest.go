package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsSignnameReportAPIRequest 服务商存量短信签名上传 API请求
// taobao.jst.sms.signname.report
//
// 用于上传目前已经在阿里通信申请到的且正在使用的签名数据，确保签名数据正确，否则会导致短信发送失败！！！
type TaobaoJstSmsSignnameReportAPIRequest struct {
	model.Params
	// 上传签名入参
	_smsSignNameRequest *SmsSignNameRequest
}

// NewTaobaoJstSmsSignnameReportRequest 初始化TaobaoJstSmsSignnameReportAPIRequest对象
func NewTaobaoJstSmsSignnameReportRequest() *TaobaoJstSmsSignnameReportAPIRequest {
	return &TaobaoJstSmsSignnameReportAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstSmsSignnameReportAPIRequest) Reset() {
	r._smsSignNameRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsSignnameReportAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.signname.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsSignnameReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstSmsSignnameReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSmsSignNameRequest is SmsSignNameRequest Setter
// 上传签名入参
func (r *TaobaoJstSmsSignnameReportAPIRequest) SetSmsSignNameRequest(_smsSignNameRequest *SmsSignNameRequest) error {
	r._smsSignNameRequest = _smsSignNameRequest
	r.Set("sms_sign_name_request", _smsSignNameRequest)
	return nil
}

// GetSmsSignNameRequest SmsSignNameRequest Getter
func (r TaobaoJstSmsSignnameReportAPIRequest) GetSmsSignNameRequest() *SmsSignNameRequest {
	return r._smsSignNameRequest
}

var poolTaobaoJstSmsSignnameReportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstSmsSignnameReportRequest()
	},
}

// GetTaobaoJstSmsSignnameReportRequest 从 sync.Pool 获取 TaobaoJstSmsSignnameReportAPIRequest
func GetTaobaoJstSmsSignnameReportAPIRequest() *TaobaoJstSmsSignnameReportAPIRequest {
	return poolTaobaoJstSmsSignnameReportAPIRequest.Get().(*TaobaoJstSmsSignnameReportAPIRequest)
}

// ReleaseTaobaoJstSmsSignnameReportAPIRequest 将 TaobaoJstSmsSignnameReportAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstSmsSignnameReportAPIRequest(v *TaobaoJstSmsSignnameReportAPIRequest) {
	v.Reset()
	poolTaobaoJstSmsSignnameReportAPIRequest.Put(v)
}
