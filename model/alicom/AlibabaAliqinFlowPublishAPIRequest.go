package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinflowpublishAPIRequest 流量发放(用户id) API请求
// alibaba.aliqin.flow.publish
//
// 阿里通信流量下发功能，允许用户补发
type AlibabaaliqinflowpublishAPIRequest struct {
	model.Params
	// 用户id
	_userId string
	// 流量
	_flow string
	// 原因
	_reason string
	// 唯一流水号（字母+数字）
	_serial string
	// 设置true为始终发送成功
	_always string
}

// NewAlibabaaliqinflowpublishRequest 初始化AlibabaaliqinflowpublishAPIRequest对象
func NewAlibabaaliqinflowpublishRequest() *AlibabaaliqinflowpublishAPIRequest {
	return &AlibabaaliqinflowpublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinflowpublishAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinflowpublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinflowpublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 用户id
func (r *AlibabaaliqinflowpublishAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaaliqinflowpublishAPIRequest) GetUserId() string {
	return r._userId
}

// SetFlow is Flow Setter
// 流量
func (r *AlibabaaliqinflowpublishAPIRequest) SetFlow(_flow string) error {
	r._flow = _flow
	r.Set("flow", _flow)
	return nil
}

// GetFlow Flow Getter
func (r AlibabaaliqinflowpublishAPIRequest) GetFlow() string {
	return r._flow
}

// SetReason is Reason Setter
// 原因
func (r *AlibabaaliqinflowpublishAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r AlibabaaliqinflowpublishAPIRequest) GetReason() string {
	return r._reason
}

// SetSerial is Serial Setter
// 唯一流水号（字母+数字）
func (r *AlibabaaliqinflowpublishAPIRequest) SetSerial(_serial string) error {
	r._serial = _serial
	r.Set("serial", _serial)
	return nil
}

// GetSerial Serial Getter
func (r AlibabaaliqinflowpublishAPIRequest) GetSerial() string {
	return r._serial
}

// SetAlways is Always Setter
// 设置true为始终发送成功
func (r *AlibabaaliqinflowpublishAPIRequest) SetAlways(_always string) error {
	r._always = _always
	r.Set("always", _always)
	return nil
}

// GetAlways Always Getter
func (r AlibabaaliqinflowpublishAPIRequest) GetAlways() string {
	return r._always
}
