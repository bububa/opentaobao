package yunosappstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外投广告上报接口 API请求
yunos.appstore.open.reportad

外投广告回流上报接口
*/
type YunosAppstoreOpenReportadRequest struct {
    model.Params
    // 广告跟踪id列表
    traceIds   []string
    // 事件时间，当前毫秒数
    time   int64
    // 事件类型：0 代表曝光事件；1 代表点击下载事件；2 代表下载完成事件；3 代表安装完成事件
    action   int64
    // 客户端版本号
    clientVerCode   int64
    // 客户端设备标识
    deviceId   string
}

// 初始化YunosAppstoreOpenReportadRequest对象
func NewYunosAppstoreOpenReportadRequest() *YunosAppstoreOpenReportadRequest{
    return &YunosAppstoreOpenReportadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosAppstoreOpenReportadRequest) GetApiMethodName() string {
    return "yunos.appstore.open.reportad"
}

// IRequest interface 方法, 获取API参数
func (r YunosAppstoreOpenReportadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TraceIds Setter
// 广告跟踪id列表
func (r *YunosAppstoreOpenReportadRequest) SetTraceIds(traceIds []string) error {
    r.traceIds = traceIds
    r.Set("trace_ids", traceIds)
    return nil
}

// TraceIds Getter
func (r YunosAppstoreOpenReportadRequest) GetTraceIds() []string {
    return r.traceIds
}
// Time Setter
// 事件时间，当前毫秒数
func (r *YunosAppstoreOpenReportadRequest) SetTime(time int64) error {
    r.time = time
    r.Set("time", time)
    return nil
}

// Time Getter
func (r YunosAppstoreOpenReportadRequest) GetTime() int64 {
    return r.time
}
// Action Setter
// 事件类型：0 代表曝光事件；1 代表点击下载事件；2 代表下载完成事件；3 代表安装完成事件
func (r *YunosAppstoreOpenReportadRequest) SetAction(action int64) error {
    r.action = action
    r.Set("action", action)
    return nil
}

// Action Getter
func (r YunosAppstoreOpenReportadRequest) GetAction() int64 {
    return r.action
}
// ClientVerCode Setter
// 客户端版本号
func (r *YunosAppstoreOpenReportadRequest) SetClientVerCode(clientVerCode int64) error {
    r.clientVerCode = clientVerCode
    r.Set("client_ver_code", clientVerCode)
    return nil
}

// ClientVerCode Getter
func (r YunosAppstoreOpenReportadRequest) GetClientVerCode() int64 {
    return r.clientVerCode
}
// DeviceId Setter
// 客户端设备标识
func (r *YunosAppstoreOpenReportadRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r YunosAppstoreOpenReportadRequest) GetDeviceId() string {
    return r.deviceId
}
