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
    appId   string
    // 传输流水号（uuid)
    processId   string
    // 事件编号（uuid）
    eventId   string
    // 事件类型，10：基础信息数据子集 ； 20：应用信息数据子集
    eventType   string
    // 事件子类型
    subType   string
    // 统一社会信用代码
    uscId   string
    // 文件内容 zip压缩+base64转码
    data   string
    // 时间戳，yyyy-MM-dd HH:mm:ss
    tiemstamp   string
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
func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

// AppId Getter
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetAppId() string {
    return r.appId
}
// ProcessId Setter
// 传输流水号（uuid)
func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetProcessId(processId string) error {
    r.processId = processId
    r.Set("process_id", processId)
    return nil
}

// ProcessId Getter
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetProcessId() string {
    return r.processId
}
// EventId Setter
// 事件编号（uuid）
func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetEventId(eventId string) error {
    r.eventId = eventId
    r.Set("event_id", eventId)
    return nil
}

// EventId Getter
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetEventId() string {
    return r.eventId
}
// EventType Setter
// 事件类型，10：基础信息数据子集 ； 20：应用信息数据子集
func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetEventType(eventType string) error {
    r.eventType = eventType
    r.Set("event_type", eventType)
    return nil
}

// EventType Getter
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetEventType() string {
    return r.eventType
}
// SubType Setter
// 事件子类型
func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetSubType(subType string) error {
    r.subType = subType
    r.Set("sub_type", subType)
    return nil
}

// SubType Getter
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetSubType() string {
    return r.subType
}
// UscId Setter
// 统一社会信用代码
func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetUscId(uscId string) error {
    r.uscId = uscId
    r.Set("usc_id", uscId)
    return nil
}

// UscId Getter
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetUscId() string {
    return r.uscId
}
// Data Setter
// 文件内容 zip压缩+base64转码
func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

// Data Getter
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetData() string {
    return r.data
}
// Tiemstamp Setter
// 时间戳，yyyy-MM-dd HH:mm:ss
func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetTiemstamp(tiemstamp string) error {
    r.tiemstamp = tiemstamp
    r.Set("tiemstamp", tiemstamp)
    return nil
}

// Tiemstamp Getter
func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetTiemstamp() string {
    return r.tiemstamp
}
