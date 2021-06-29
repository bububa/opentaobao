package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
协同平台数据下行接口 APIRequest
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

func NewAlibabaCfdaXtptAppAcceptInfoRequest() *AlibabaCfdaXtptAppAcceptInfoRequest{
    return &AlibabaCfdaXtptAppAcceptInfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetApiMethodName() string {
    return "alibaba.cfda.xtpt.app.accept.info"
}

func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetAppId() string {
    return r.appId
}

func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetProcessId(processId string) error {
    r.processId = processId
    r.Set("process_id", processId)
    return nil
}

func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetProcessId() string {
    return r.processId
}

func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetEventId(eventId string) error {
    r.eventId = eventId
    r.Set("event_id", eventId)
    return nil
}

func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetEventId() string {
    return r.eventId
}

func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetEventType(eventType string) error {
    r.eventType = eventType
    r.Set("event_type", eventType)
    return nil
}

func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetEventType() string {
    return r.eventType
}

func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetSubType(subType string) error {
    r.subType = subType
    r.Set("sub_type", subType)
    return nil
}

func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetSubType() string {
    return r.subType
}

func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetUscId(uscId string) error {
    r.uscId = uscId
    r.Set("usc_id", uscId)
    return nil
}

func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetUscId() string {
    return r.uscId
}

func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetData() string {
    return r.data
}

func (r *AlibabaCfdaXtptAppAcceptInfoRequest) SetTiemstamp(tiemstamp string) error {
    r.tiemstamp = tiemstamp
    r.Set("tiemstamp", tiemstamp)
    return nil
}

func (r AlibabaCfdaXtptAppAcceptInfoRequest) GetTiemstamp() string {
    return r.tiemstamp
}

