package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商更新预警消息状态 API请求
tmall.servicecenter.servicemonitormessage.update

服务商收到预警后，需要进行回复已读状态，并可填写备注
*/
type TmallServicecenterServicemonitormessageUpdateRequest struct {
    model.Params
    // 预警消息id
    _serviceMonitorMessageId   int64
    // 预警处理备注
    _memo   string
    // 可更新状态：3、已读
    _status   int64
}

// 初始化TmallServicecenterServicemonitormessageUpdateRequest对象
func NewTmallServicecenterServicemonitormessageUpdateRequest() *TmallServicecenterServicemonitormessageUpdateRequest{
    return &TmallServicecenterServicemonitormessageUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicemonitormessageUpdateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicemonitormessage.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicemonitormessageUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceMonitorMessageId Setter
// 预警消息id
func (r *TmallServicecenterServicemonitormessageUpdateRequest) SetServiceMonitorMessageId(_serviceMonitorMessageId int64) error {
    r._serviceMonitorMessageId = _serviceMonitorMessageId
    r.Set("service_monitor_message_id", _serviceMonitorMessageId)
    return nil
}

// ServiceMonitorMessageId Getter
func (r TmallServicecenterServicemonitormessageUpdateRequest) GetServiceMonitorMessageId() int64 {
    return r._serviceMonitorMessageId
}
// Memo Setter
// 预警处理备注
func (r *TmallServicecenterServicemonitormessageUpdateRequest) SetMemo(_memo string) error {
    r._memo = _memo
    r.Set("memo", _memo)
    return nil
}

// Memo Getter
func (r TmallServicecenterServicemonitormessageUpdateRequest) GetMemo() string {
    return r._memo
}
// Status Setter
// 可更新状态：3、已读
func (r *TmallServicecenterServicemonitormessageUpdateRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TmallServicecenterServicemonitormessageUpdateRequest) GetStatus() int64 {
    return r._status
}
