package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCfdaXtptAppAcceptInfoAPIRequest 协同平台数据下行接口 API请求
// alibaba.cfda.xtpt.app.accept.info
//
// 协同平台数据下行接口
type AlibabaCfdaXtptAppAcceptInfoAPIRequest struct {
	model.Params
	// 接入系统唯一标识，由协同平台分配
	_appId string
	// 传输流水号（uuid)
	_processId string
	// 事件编号（uuid）
	_eventId string
	// 事件类型，10：基础信息数据子集 ； 20：应用信息数据子集
	_eventType string
	// 事件子类型
	_subType string
	// 统一社会信用代码
	_uscId string
	// 文件内容 zip压缩+base64转码
	_data string
	// 时间戳，yyyy-MM-dd HH:mm:ss
	_tiemstamp string
}

// NewAlibabaCfdaXtptAppAcceptInfoRequest 初始化AlibabaCfdaXtptAppAcceptInfoAPIRequest对象
func NewAlibabaCfdaXtptAppAcceptInfoRequest() *AlibabaCfdaXtptAppAcceptInfoAPIRequest {
	return &AlibabaCfdaXtptAppAcceptInfoAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) Reset() {
	r._appId = ""
	r._processId = ""
	r._eventId = ""
	r._eventType = ""
	r._subType = ""
	r._uscId = ""
	r._data = ""
	r._tiemstamp = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetApiMethodName() string {
	return "alibaba.cfda.xtpt.app.accept.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppId is AppId Setter
// 接入系统唯一标识，由协同平台分配
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetAppId() string {
	return r._appId
}

// SetProcessId is ProcessId Setter
// 传输流水号（uuid)
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) SetProcessId(_processId string) error {
	r._processId = _processId
	r.Set("process_id", _processId)
	return nil
}

// GetProcessId ProcessId Getter
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetProcessId() string {
	return r._processId
}

// SetEventId is EventId Setter
// 事件编号（uuid）
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) SetEventId(_eventId string) error {
	r._eventId = _eventId
	r.Set("event_id", _eventId)
	return nil
}

// GetEventId EventId Getter
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetEventId() string {
	return r._eventId
}

// SetEventType is EventType Setter
// 事件类型，10：基础信息数据子集 ； 20：应用信息数据子集
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) SetEventType(_eventType string) error {
	r._eventType = _eventType
	r.Set("event_type", _eventType)
	return nil
}

// GetEventType EventType Getter
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetEventType() string {
	return r._eventType
}

// SetSubType is SubType Setter
// 事件子类型
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) SetSubType(_subType string) error {
	r._subType = _subType
	r.Set("sub_type", _subType)
	return nil
}

// GetSubType SubType Getter
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetSubType() string {
	return r._subType
}

// SetUscId is UscId Setter
// 统一社会信用代码
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) SetUscId(_uscId string) error {
	r._uscId = _uscId
	r.Set("usc_id", _uscId)
	return nil
}

// GetUscId UscId Getter
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetUscId() string {
	return r._uscId
}

// SetData is Data Setter
// 文件内容 zip压缩+base64转码
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetData() string {
	return r._data
}

// SetTiemstamp is Tiemstamp Setter
// 时间戳，yyyy-MM-dd HH:mm:ss
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) SetTiemstamp(_tiemstamp string) error {
	r._tiemstamp = _tiemstamp
	r.Set("tiemstamp", _tiemstamp)
	return nil
}

// GetTiemstamp Tiemstamp Getter
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetTiemstamp() string {
	return r._tiemstamp
}

var poolAlibabaCfdaXtptAppAcceptInfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCfdaXtptAppAcceptInfoRequest()
	},
}

// GetAlibabaCfdaXtptAppAcceptInfoRequest 从 sync.Pool 获取 AlibabaCfdaXtptAppAcceptInfoAPIRequest
func GetAlibabaCfdaXtptAppAcceptInfoAPIRequest() *AlibabaCfdaXtptAppAcceptInfoAPIRequest {
	return poolAlibabaCfdaXtptAppAcceptInfoAPIRequest.Get().(*AlibabaCfdaXtptAppAcceptInfoAPIRequest)
}

// ReleaseAlibabaCfdaXtptAppAcceptInfoAPIRequest 将 AlibabaCfdaXtptAppAcceptInfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaCfdaXtptAppAcceptInfoAPIRequest(v *AlibabaCfdaXtptAppAcceptInfoAPIRequest) {
	v.Reset()
	poolAlibabaCfdaXtptAppAcceptInfoAPIRequest.Put(v)
}
