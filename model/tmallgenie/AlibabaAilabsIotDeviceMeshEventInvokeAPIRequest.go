package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest 弹内设备中心事件调用 API请求
// alibaba.ailabs.iot.device.mesh.event.invoke
//
// 弹内设备中心事件调用
type AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest struct {
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

// NewAlibabaAilabsIotDeviceMeshEventInvokeRequest 初始化AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest对象
func NewAlibabaAilabsIotDeviceMeshEventInvokeRequest() *AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest {
	return &AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.device.mesh.event.invoke"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceId is TraceId Setter
// 链接id
func (r *AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetUserType is UserType Setter
// 用户类型
func (r *AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// GetUserType UserType Getter
func (r AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest) GetUserType() string {
	return r._userType
}

// SetServerEventRequestId is ServerEventRequestId Setter
// 事件服务请求id
func (r *AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest) SetServerEventRequestId(_serverEventRequestId string) error {
	r._serverEventRequestId = _serverEventRequestId
	r.Set("server_event_request_id", _serverEventRequestId)
	return nil
}

// GetServerEventRequestId ServerEventRequestId Getter
func (r AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest) GetServerEventRequestId() string {
	return r._serverEventRequestId
}

// SetUuid is Uuid Setter
// 音箱uuid
func (r *AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest) GetUuid() string {
	return r._uuid
}

// SetEvent is Event Setter
// 事件
func (r *AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest) SetEvent(_event *LivingGenericEventDto) error {
	r._event = _event
	r.Set("event", _event)
	return nil
}

// GetEvent Event Getter
func (r AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest) GetEvent() *LivingGenericEventDto {
	return r._event
}

// SetUserId is UserId Setter
// 用户
func (r *AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest) GetUserId() int64 {
	return r._userId
}
