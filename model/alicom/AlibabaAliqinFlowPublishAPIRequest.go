package alicom

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFlowPublishAPIRequest) Reset() {
	r._userId = ""
	r._flow = ""
	r._reason = ""
	r._serial = ""
	r._always = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowPublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFlowPublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 用户id
func (r *AlibabaAliqinFlowPublishAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaAliqinFlowPublishAPIRequest) GetUserId() string {
	return r._userId
}

// SetFlow is Flow Setter
// 流量
func (r *AlibabaAliqinFlowPublishAPIRequest) SetFlow(_flow string) error {
	r._flow = _flow
	r.Set("flow", _flow)
	return nil
}

// GetFlow Flow Getter
func (r AlibabaAliqinFlowPublishAPIRequest) GetFlow() string {
	return r._flow
}

// SetReason is Reason Setter
// 原因
func (r *AlibabaAliqinFlowPublishAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r AlibabaAliqinFlowPublishAPIRequest) GetReason() string {
	return r._reason
}

// SetSerial is Serial Setter
// 唯一流水号（字母+数字）
func (r *AlibabaAliqinFlowPublishAPIRequest) SetSerial(_serial string) error {
	r._serial = _serial
	r.Set("serial", _serial)
	return nil
}

// GetSerial Serial Getter
func (r AlibabaAliqinFlowPublishAPIRequest) GetSerial() string {
	return r._serial
}

// SetAlways is Always Setter
// 设置true为始终发送成功
func (r *AlibabaAliqinFlowPublishAPIRequest) SetAlways(_always string) error {
	r._always = _always
	r.Set("always", _always)
	return nil
}

// GetAlways Always Getter
func (r AlibabaAliqinFlowPublishAPIRequest) GetAlways() string {
	return r._always
}

var poolAlibabaAliqinFlowPublishAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFlowPublishRequest()
	},
}

// GetAlibabaAliqinFlowPublishRequest 从 sync.Pool 获取 AlibabaAliqinFlowPublishAPIRequest
func GetAlibabaAliqinFlowPublishAPIRequest() *AlibabaAliqinFlowPublishAPIRequest {
	return poolAlibabaAliqinFlowPublishAPIRequest.Get().(*AlibabaAliqinFlowPublishAPIRequest)
}

// ReleaseAlibabaAliqinFlowPublishAPIRequest 将 AlibabaAliqinFlowPublishAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFlowPublishAPIRequest(v *AlibabaAliqinFlowPublishAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFlowPublishAPIRequest.Put(v)
}
