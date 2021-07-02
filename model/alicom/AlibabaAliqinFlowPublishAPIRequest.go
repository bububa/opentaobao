package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowPublishAPIRequest 流量发放(用户id) API请求
// alibaba.aliqin.flow.publish
//
// 阿里通信流量下发功能，允许用户补发
type AlibabaAliqinFlowPublishAPIRequest struct {
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

// NewAlibabaAliqinFlowPublishRequest 初始化AlibabaAliqinFlowPublishAPIRequest对象
func NewAlibabaAliqinFlowPublishRequest() *AlibabaAliqinFlowPublishAPIRequest {
	return &AlibabaAliqinFlowPublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowPublishAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UserId Setter
// 用户id
func (r *AlibabaAliqinFlowPublishAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r AlibabaAliqinFlowPublishAPIRequest) GetUserId() string {
	return r._userId
}

// Set is Flow Setter
// 流量
func (r *AlibabaAliqinFlowPublishAPIRequest) SetFlow(_flow string) error {
	r._flow = _flow
	r.Set("flow", _flow)
	return nil
}

// Get Flow Getter
func (r AlibabaAliqinFlowPublishAPIRequest) GetFlow() string {
	return r._flow
}

// Set is Reason Setter
// 原因
func (r *AlibabaAliqinFlowPublishAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// Get Reason Getter
func (r AlibabaAliqinFlowPublishAPIRequest) GetReason() string {
	return r._reason
}

// Set is Serial Setter
// 唯一流水号（字母+数字）
func (r *AlibabaAliqinFlowPublishAPIRequest) SetSerial(_serial string) error {
	r._serial = _serial
	r.Set("serial", _serial)
	return nil
}

// Get Serial Getter
func (r AlibabaAliqinFlowPublishAPIRequest) GetSerial() string {
	return r._serial
}

// Set is Always Setter
// 设置true为始终发送成功
func (r *AlibabaAliqinFlowPublishAPIRequest) SetAlways(_always string) error {
	r._always = _always
	r.Set("always", _always)
	return nil
}

// Get Always Getter
func (r AlibabaAliqinFlowPublishAPIRequest) GetAlways() string {
	return r._always
}
