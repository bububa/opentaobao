package yunosappstore

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosAppstoreOpenReportadAPIRequest 外投广告上报接口 API请求
// yunos.appstore.open.reportad
//
// 外投广告回流上报接口
type YunosAppstoreOpenReportadAPIRequest struct {
	model.Params
	// 广告跟踪id列表
	_traceIds []string
	// 客户端设备标识
	_deviceId string
	// 事件时间，当前毫秒数
	_time int64
	// 事件类型：0 代表曝光事件；1 代表点击下载事件；2 代表下载完成事件；3 代表安装完成事件
	_action int64
	// 客户端版本号
	_clientVerCode int64
}

// NewYunosAppstoreOpenReportadRequest 初始化YunosAppstoreOpenReportadAPIRequest对象
func NewYunosAppstoreOpenReportadRequest() *YunosAppstoreOpenReportadAPIRequest {
	return &YunosAppstoreOpenReportadAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosAppstoreOpenReportadAPIRequest) Reset() {
	r._traceIds = r._traceIds[:0]
	r._deviceId = ""
	r._time = 0
	r._action = 0
	r._clientVerCode = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosAppstoreOpenReportadAPIRequest) GetApiMethodName() string {
	return "yunos.appstore.open.reportad"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosAppstoreOpenReportadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosAppstoreOpenReportadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceIds is TraceIds Setter
// 广告跟踪id列表
func (r *YunosAppstoreOpenReportadAPIRequest) SetTraceIds(_traceIds []string) error {
	r._traceIds = _traceIds
	r.Set("trace_ids", _traceIds)
	return nil
}

// GetTraceIds TraceIds Getter
func (r YunosAppstoreOpenReportadAPIRequest) GetTraceIds() []string {
	return r._traceIds
}

// SetDeviceId is DeviceId Setter
// 客户端设备标识
func (r *YunosAppstoreOpenReportadAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r YunosAppstoreOpenReportadAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetTime is Time Setter
// 事件时间，当前毫秒数
func (r *YunosAppstoreOpenReportadAPIRequest) SetTime(_time int64) error {
	r._time = _time
	r.Set("time", _time)
	return nil
}

// GetTime Time Getter
func (r YunosAppstoreOpenReportadAPIRequest) GetTime() int64 {
	return r._time
}

// SetAction is Action Setter
// 事件类型：0 代表曝光事件；1 代表点击下载事件；2 代表下载完成事件；3 代表安装完成事件
func (r *YunosAppstoreOpenReportadAPIRequest) SetAction(_action int64) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r YunosAppstoreOpenReportadAPIRequest) GetAction() int64 {
	return r._action
}

// SetClientVerCode is ClientVerCode Setter
// 客户端版本号
func (r *YunosAppstoreOpenReportadAPIRequest) SetClientVerCode(_clientVerCode int64) error {
	r._clientVerCode = _clientVerCode
	r.Set("client_ver_code", _clientVerCode)
	return nil
}

// GetClientVerCode ClientVerCode Getter
func (r YunosAppstoreOpenReportadAPIRequest) GetClientVerCode() int64 {
	return r._clientVerCode
}

var poolYunosAppstoreOpenReportadAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosAppstoreOpenReportadRequest()
	},
}

// GetYunosAppstoreOpenReportadRequest 从 sync.Pool 获取 YunosAppstoreOpenReportadAPIRequest
func GetYunosAppstoreOpenReportadAPIRequest() *YunosAppstoreOpenReportadAPIRequest {
	return poolYunosAppstoreOpenReportadAPIRequest.Get().(*YunosAppstoreOpenReportadAPIRequest)
}

// ReleaseYunosAppstoreOpenReportadAPIRequest 将 YunosAppstoreOpenReportadAPIRequest 放入 sync.Pool
func ReleaseYunosAppstoreOpenReportadAPIRequest(v *YunosAppstoreOpenReportadAPIRequest) {
	v.Reset()
	poolYunosAppstoreOpenReportadAPIRequest.Put(v)
}
