package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商更新预警消息状态 APIRequest
tmall.servicecenter.servicemonitormessage.update

服务商收到预警后，需要进行回复已读状态，并可填写备注
*/
type TmallServicecenterServicemonitormessageUpdateRequest struct {
    model.Params

    // 预警消息id
    serviceMonitorMessageId   int64 

    // 预警处理备注
    memo   string 

    // 可更新状态：3、已读
    status   int64 

}

func NewTmallServicecenterServicemonitormessageUpdateRequest() *TmallServicecenterServicemonitormessageUpdateRequest{
    return &TmallServicecenterServicemonitormessageUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterServicemonitormessageUpdateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicemonitormessage.update"
}

func (r TmallServicecenterServicemonitormessageUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterServicemonitormessageUpdateRequest) SetServiceMonitorMessageId(serviceMonitorMessageId int64) error {
    r.serviceMonitorMessageId = serviceMonitorMessageId
    r.Set("service_monitor_message_id", serviceMonitorMessageId)
    return nil
}

func (r TmallServicecenterServicemonitormessageUpdateRequest) GetServiceMonitorMessageId() int64 {
    return r.serviceMonitorMessageId
}

func (r *TmallServicecenterServicemonitormessageUpdateRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

func (r TmallServicecenterServicemonitormessageUpdateRequest) GetMemo() string {
    return r.memo
}

func (r *TmallServicecenterServicemonitormessageUpdateRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TmallServicecenterServicemonitormessageUpdateRequest) GetStatus() int64 {
    return r.status
}

