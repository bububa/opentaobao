package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
协同平台数据下行接口 API请求
alibaba.cfda.xtpt.app.accept.info

协同平台数据下行接口
*/
type AlibabaCfdaXtptAppAcceptInfoRequest struct {
    model.Params
    // 接入系统唯一标识，由协同平台分配
    _appId   string
    // 传输流水号（uuid)
    _processId   string
    // 事件编号（uuid）
    _eventId   string
    // 事件类型，10：基础信息数据子集 ； 20：应用信息数据子集
    _eventType   string
    // 事件子类型
    _subType   string
    // 统一社会信用代码
    _uscId   string
    // 文件内容 zip压缩+base64转码
    _data   string
    // 时间戳，yyyy-MM-dd HH:mm:ss
    _tiemstamp   string
}

// 初始化AlibabaCfdaXtptAppAcceptInfoRequest对象
func NewAlibabaCfdaXtptAppAcceptInfoRequest() *AlibabaCfdaXtptAppAcceptInfoRequest{
    return &AlibabaCfdaXtptAppAcceptInfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetApiMethodName() string {
    return "alibaba.cfda.xtpt.app.accept.info"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppId Setter
// 接入系统唯一标识，由协同平台分配
func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetAppId(_appId string) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetAppId() string {
    return r._appId
}
// ProcessId Setter
// 传输流水号（uuid)
func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetProcessId(_processId string) error {
    r._processId = _processId
    r.Set("process_id", _processId)
    return nil
}

// ProcessId Getter
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetProcessId() string {
    return r._processId
}
// EventId Setter
// 事件编号（uuid）
func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetEventId(_eventId string) error {
    r._eventId = _eventId
    r.Set("event_id", _eventId)
    return nil
}

// EventId Getter
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetEventId() string {
    return r._eventId
}
// EventType Setter
// 事件类型，10：基础信息数据子集 ； 20：应用信息数据子集
func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetEventType(_eventType string) error {
    r._eventType = _eventType
    r.Set("event_type", _eventType)
    return nil
}

// EventType Getter
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetEventType() string {
    return r._eventType
}
// SubType Setter
// 事件子类型
func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetSubType(_subType string) error {
    r._subType = _subType
    r.Set("sub_type", _subType)
    return nil
}

// SubType Getter
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetSubType() string {
    return r._subType
}
// UscId Setter
// 统一社会信用代码
func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetUscId(_uscId string) error {
    r._uscId = _uscId
    r.Set("usc_id", _uscId)
    return nil
}

// UscId Getter
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetUscId() string {
    return r._uscId
}
// Data Setter
// 文件内容 zip压缩+base64转码
func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetData(_data string) error {
    r._data = _data
    r.Set("data", _data)
    return nil
}

// Data Getter
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetData() string {
    return r._data
}
// Tiemstamp Setter
// 时间戳，yyyy-MM-dd HH:mm:ss
func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetTiemstamp(_tiemstamp string) error {
    r._tiemstamp = _tiemstamp
    r.Set("tiemstamp", _tiemstamp)
    return nil
}

// Tiemstamp Getter
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetTiemstamp() string {
    return r._tiemstamp
}
