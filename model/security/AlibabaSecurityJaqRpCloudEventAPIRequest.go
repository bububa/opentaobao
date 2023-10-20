package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpCloudEventAPIRequest 事件上报 API请求
// alibaba.security.jaq.rp.cloud.event
//
// 事件上报接口
type AlibabaSecurityJaqRpCloudEventAPIRequest struct {
	model.Params
	// 认证token
	_verifyToken string
	// 事件编码
	_eventCode string
	// 事件信息
	_eventData string
}

// NewAlibabaSecurityJaqRpCloudEventRequest 初始化AlibabaSecurityJaqRpCloudEventAPIRequest对象
func NewAlibabaSecurityJaqRpCloudEventRequest() *AlibabaSecurityJaqRpCloudEventAPIRequest {
	return &AlibabaSecurityJaqRpCloudEventAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqRpCloudEventAPIRequest) Reset() {
	r._verifyToken = ""
	r._eventCode = ""
	r._eventData = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudEventAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.cloud.event"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudEventAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqRpCloudEventAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// 认证token
func (r *AlibabaSecurityJaqRpCloudEventAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabaSecurityJaqRpCloudEventAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

// SetEventCode is EventCode Setter
// 事件编码
func (r *AlibabaSecurityJaqRpCloudEventAPIRequest) SetEventCode(_eventCode string) error {
	r._eventCode = _eventCode
	r.Set("event_code", _eventCode)
	return nil
}

// GetEventCode EventCode Getter
func (r AlibabaSecurityJaqRpCloudEventAPIRequest) GetEventCode() string {
	return r._eventCode
}

// SetEventData is EventData Setter
// 事件信息
func (r *AlibabaSecurityJaqRpCloudEventAPIRequest) SetEventData(_eventData string) error {
	r._eventData = _eventData
	r.Set("event_data", _eventData)
	return nil
}

// GetEventData EventData Getter
func (r AlibabaSecurityJaqRpCloudEventAPIRequest) GetEventData() string {
	return r._eventData
}

var poolAlibabaSecurityJaqRpCloudEventAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqRpCloudEventRequest()
	},
}

// GetAlibabaSecurityJaqRpCloudEventRequest 从 sync.Pool 获取 AlibabaSecurityJaqRpCloudEventAPIRequest
func GetAlibabaSecurityJaqRpCloudEventAPIRequest() *AlibabaSecurityJaqRpCloudEventAPIRequest {
	return poolAlibabaSecurityJaqRpCloudEventAPIRequest.Get().(*AlibabaSecurityJaqRpCloudEventAPIRequest)
}

// ReleaseAlibabaSecurityJaqRpCloudEventAPIRequest 将 AlibabaSecurityJaqRpCloudEventAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqRpCloudEventAPIRequest(v *AlibabaSecurityJaqRpCloudEventAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqRpCloudEventAPIRequest.Put(v)
}
