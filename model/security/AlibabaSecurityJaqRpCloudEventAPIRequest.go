package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrpcloudeventAPIRequest 事件上报 API请求
// alibaba.security.jaq.rp.cloud.event
//
// 事件上报接口
type AlibabasecurityjaqrpcloudeventAPIRequest struct {
	model.Params
	// 认证token
	_verifyToken string
	// 事件编码
	_eventCode string
	// 事件信息
	_eventData string
}

// NewAlibabasecurityjaqrpcloudeventRequest 初始化AlibabasecurityjaqrpcloudeventAPIRequest对象
func NewAlibabasecurityjaqrpcloudeventRequest() *AlibabasecurityjaqrpcloudeventAPIRequest {
	return &AlibabasecurityjaqrpcloudeventAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqrpcloudeventAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.cloud.event"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqrpcloudeventAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqrpcloudeventAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// 认证token
func (r *AlibabasecurityjaqrpcloudeventAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabasecurityjaqrpcloudeventAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

// SetEventCode is EventCode Setter
// 事件编码
func (r *AlibabasecurityjaqrpcloudeventAPIRequest) SetEventCode(_eventCode string) error {
	r._eventCode = _eventCode
	r.Set("event_code", _eventCode)
	return nil
}

// GetEventCode EventCode Getter
func (r AlibabasecurityjaqrpcloudeventAPIRequest) GetEventCode() string {
	return r._eventCode
}

// SetEventData is EventData Setter
// 事件信息
func (r *AlibabasecurityjaqrpcloudeventAPIRequest) SetEventData(_eventData string) error {
	r._eventData = _eventData
	r.Set("event_data", _eventData)
	return nil
}

// GetEventData EventData Getter
func (r AlibabasecurityjaqrpcloudeventAPIRequest) GetEventData() string {
	return r._eventData
}
