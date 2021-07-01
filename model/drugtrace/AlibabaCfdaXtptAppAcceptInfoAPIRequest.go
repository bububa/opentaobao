package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCfdaXtptAppAcceptInfoAPIRequest
协同平台数据下行接口 API请求
alibaba.cfda.xtpt.app.accept.info

协同平台数据下行接口 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetApiMethodName() string {
	return "alibaba.cfda.xtpt.app.accept.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AppId Setter
// 接入系统唯一标识，由协同平台分配
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// Get AppId Getter
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetAppId() string {
	return r._appId
}

// Set is ProcessId Setter
// 传输流水号（uuid)
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) SetProcessId(_processId string) error {
	r._processId = _processId
	r.Set("process_id", _processId)
	return nil
}

// Get ProcessId Getter
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetProcessId() string {
	return r._processId
}

// Set is EventId Setter
// 事件编号（uuid）
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) SetEventId(_eventId string) error {
	r._eventId = _eventId
	r.Set("event_id", _eventId)
	return nil
}

// Get EventId Getter
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetEventId() string {
	return r._eventId
}

// Set is EventType Setter
// 事件类型，10：基础信息数据子集 ； 20：应用信息数据子集
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) SetEventType(_eventType string) error {
	r._eventType = _eventType
	r.Set("event_type", _eventType)
	return nil
}

// Get EventType Getter
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetEventType() string {
	return r._eventType
}

// Set is SubType Setter
// 事件子类型
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) SetSubType(_subType string) error {
	r._subType = _subType
	r.Set("sub_type", _subType)
	return nil
}

// Get SubType Getter
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetSubType() string {
	return r._subType
}

// Set is UscId Setter
// 统一社会信用代码
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) SetUscId(_uscId string) error {
	r._uscId = _uscId
	r.Set("usc_id", _uscId)
	return nil
}

// Get UscId Getter
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetUscId() string {
	return r._uscId
}

// Set is Data Setter
// 文件内容 zip压缩+base64转码
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// Get Data Getter
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetData() string {
	return r._data
}

// Set is Tiemstamp Setter
// 时间戳，yyyy-MM-dd HH:mm:ss
func (r *AlibabaCfdaXtptAppAcceptInfoAPIRequest) SetTiemstamp(_tiemstamp string) error {
	r._tiemstamp = _tiemstamp
	r.Set("tiemstamp", _tiemstamp)
	return nil
}

// Get Tiemstamp Getter
func (r AlibabaCfdaXtptAppAcceptInfoAPIRequest) GetTiemstamp() string {
	return r._tiemstamp
}
