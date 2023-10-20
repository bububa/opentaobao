package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacfdaxtptappacceptinfoAPIRequest 协同平台数据下行接口 API请求
// alibaba.cfda.xtpt.app.accept.info
//
// 协同平台数据下行接口
type AlibabacfdaxtptappacceptinfoAPIRequest struct {
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

// NewAlibabacfdaxtptappacceptinfoRequest 初始化AlibabacfdaxtptappacceptinfoAPIRequest对象
func NewAlibabacfdaxtptappacceptinfoRequest() *AlibabacfdaxtptappacceptinfoAPIRequest {
	return &AlibabacfdaxtptappacceptinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacfdaxtptappacceptinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.cfda.xtpt.app.accept.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacfdaxtptappacceptinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacfdaxtptappacceptinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppId is AppId Setter
// 接入系统唯一标识，由协同平台分配
func (r *AlibabacfdaxtptappacceptinfoAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabacfdaxtptappacceptinfoAPIRequest) GetAppId() string {
	return r._appId
}

// SetProcessId is ProcessId Setter
// 传输流水号（uuid)
func (r *AlibabacfdaxtptappacceptinfoAPIRequest) SetProcessId(_processId string) error {
	r._processId = _processId
	r.Set("process_id", _processId)
	return nil
}

// GetProcessId ProcessId Getter
func (r AlibabacfdaxtptappacceptinfoAPIRequest) GetProcessId() string {
	return r._processId
}

// SetEventId is EventId Setter
// 事件编号（uuid）
func (r *AlibabacfdaxtptappacceptinfoAPIRequest) SetEventId(_eventId string) error {
	r._eventId = _eventId
	r.Set("event_id", _eventId)
	return nil
}

// GetEventId EventId Getter
func (r AlibabacfdaxtptappacceptinfoAPIRequest) GetEventId() string {
	return r._eventId
}

// SetEventType is EventType Setter
// 事件类型，10：基础信息数据子集 ； 20：应用信息数据子集
func (r *AlibabacfdaxtptappacceptinfoAPIRequest) SetEventType(_eventType string) error {
	r._eventType = _eventType
	r.Set("event_type", _eventType)
	return nil
}

// GetEventType EventType Getter
func (r AlibabacfdaxtptappacceptinfoAPIRequest) GetEventType() string {
	return r._eventType
}

// SetSubType is SubType Setter
// 事件子类型
func (r *AlibabacfdaxtptappacceptinfoAPIRequest) SetSubType(_subType string) error {
	r._subType = _subType
	r.Set("sub_type", _subType)
	return nil
}

// GetSubType SubType Getter
func (r AlibabacfdaxtptappacceptinfoAPIRequest) GetSubType() string {
	return r._subType
}

// SetUscId is UscId Setter
// 统一社会信用代码
func (r *AlibabacfdaxtptappacceptinfoAPIRequest) SetUscId(_uscId string) error {
	r._uscId = _uscId
	r.Set("usc_id", _uscId)
	return nil
}

// GetUscId UscId Getter
func (r AlibabacfdaxtptappacceptinfoAPIRequest) GetUscId() string {
	return r._uscId
}

// SetData is Data Setter
// 文件内容 zip压缩+base64转码
func (r *AlibabacfdaxtptappacceptinfoAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabacfdaxtptappacceptinfoAPIRequest) GetData() string {
	return r._data
}

// SetTiemstamp is Tiemstamp Setter
// 时间戳，yyyy-MM-dd HH:mm:ss
func (r *AlibabacfdaxtptappacceptinfoAPIRequest) SetTiemstamp(_tiemstamp string) error {
	r._tiemstamp = _tiemstamp
	r.Set("tiemstamp", _tiemstamp)
	return nil
}

// GetTiemstamp Tiemstamp Getter
func (r AlibabacfdaxtptappacceptinfoAPIRequest) GetTiemstamp() string {
	return r._tiemstamp
}
