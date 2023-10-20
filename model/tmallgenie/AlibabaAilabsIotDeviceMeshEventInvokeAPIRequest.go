package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsiotdevicemesheventinvokeAPIRequest 弹内设备中心事件调用 API请求
// alibaba.ailabs.iot.device.mesh.event.invoke
//
// 弹内设备中心事件调用
type AlibabaailabsiotdevicemesheventinvokeAPIRequest struct {
	model.Params
	// 链接id
	_traceId string
	// 用户类型
	_userType string
	// 事件服务请求id
	_serverEventRequestId string
	// 音箱uuid
	_uuid string
	// 事件
	_event *LivingGenericEventDto
	// 用户
	_userId int64
}

// NewAlibabaailabsiotdevicemesheventinvokeRequest 初始化AlibabaailabsiotdevicemesheventinvokeAPIRequest对象
func NewAlibabaailabsiotdevicemesheventinvokeRequest() *AlibabaailabsiotdevicemesheventinvokeAPIRequest {
	return &AlibabaailabsiotdevicemesheventinvokeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsiotdevicemesheventinvokeAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.device.mesh.event.invoke"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsiotdevicemesheventinvokeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsiotdevicemesheventinvokeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceId is TraceId Setter
// 链接id
func (r *AlibabaailabsiotdevicemesheventinvokeAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r AlibabaailabsiotdevicemesheventinvokeAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetUserType is UserType Setter
// 用户类型
func (r *AlibabaailabsiotdevicemesheventinvokeAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// GetUserType UserType Getter
func (r AlibabaailabsiotdevicemesheventinvokeAPIRequest) GetUserType() string {
	return r._userType
}

// SetServerEventRequestId is ServerEventRequestId Setter
// 事件服务请求id
func (r *AlibabaailabsiotdevicemesheventinvokeAPIRequest) SetServerEventRequestId(_serverEventRequestId string) error {
	r._serverEventRequestId = _serverEventRequestId
	r.Set("server_event_request_id", _serverEventRequestId)
	return nil
}

// GetServerEventRequestId ServerEventRequestId Getter
func (r AlibabaailabsiotdevicemesheventinvokeAPIRequest) GetServerEventRequestId() string {
	return r._serverEventRequestId
}

// SetUuid is Uuid Setter
// 音箱uuid
func (r *AlibabaailabsiotdevicemesheventinvokeAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaailabsiotdevicemesheventinvokeAPIRequest) GetUuid() string {
	return r._uuid
}

// SetEvent is Event Setter
// 事件
func (r *AlibabaailabsiotdevicemesheventinvokeAPIRequest) SetEvent(_event *LivingGenericEventDto) error {
	r._event = _event
	r.Set("event", _event)
	return nil
}

// GetEvent Event Getter
func (r AlibabaailabsiotdevicemesheventinvokeAPIRequest) GetEvent() *LivingGenericEventDto {
	return r._event
}

// SetUserId is UserId Setter
// 用户
func (r *AlibabaailabsiotdevicemesheventinvokeAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaailabsiotdevicemesheventinvokeAPIRequest) GetUserId() int64 {
	return r._userId
}
