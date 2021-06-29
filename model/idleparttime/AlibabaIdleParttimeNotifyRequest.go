package idleparttime

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
兼职通知接口 APIRequest
alibaba.idle.parttime.notify

服务商侧有岗位状态变更对我们进行通知(岗位关闭, 录取状态)
*/
type AlibabaIdleParttimeNotifyRequest struct {
    model.Params

    // 实时同步类型, 0: 岗位状态, 1: 录取状态
    type   int64 

    // 岗位: 0关闭 ; 录取: 0不录取, 1录取
    status   int64 

    // 岗位id
    jobId   int64 

    // 用户id
    userId   int64 

    // 同步时间
    syncTime   int64 

    // 报名id
    applyId   int64 

    // 通知消息
    message   string 

}

func NewAlibabaIdleParttimeNotifyRequest() *AlibabaIdleParttimeNotifyRequest{
    return &AlibabaIdleParttimeNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleParttimeNotifyRequest) GetApiMethodName() string {
    return "alibaba.idle.parttime.notify"
}

func (r AlibabaIdleParttimeNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleParttimeNotifyRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaIdleParttimeNotifyRequest) GetType() int64 {
    return r.type
}

func (r *AlibabaIdleParttimeNotifyRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaIdleParttimeNotifyRequest) GetStatus() int64 {
    return r.status
}

func (r *AlibabaIdleParttimeNotifyRequest) SetJobId(jobId int64) error {
    r.jobId = jobId
    r.Set("job_id", jobId)
    return nil
}

func (r AlibabaIdleParttimeNotifyRequest) GetJobId() int64 {
    return r.jobId
}

func (r *AlibabaIdleParttimeNotifyRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaIdleParttimeNotifyRequest) GetUserId() int64 {
    return r.userId
}

func (r *AlibabaIdleParttimeNotifyRequest) SetSyncTime(syncTime int64) error {
    r.syncTime = syncTime
    r.Set("sync_time", syncTime)
    return nil
}

func (r AlibabaIdleParttimeNotifyRequest) GetSyncTime() int64 {
    return r.syncTime
}

func (r *AlibabaIdleParttimeNotifyRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r AlibabaIdleParttimeNotifyRequest) GetApplyId() int64 {
    return r.applyId
}

func (r *AlibabaIdleParttimeNotifyRequest) SetMessage(message string) error {
    r.message = message
    r.Set("message", message)
    return nil
}

func (r AlibabaIdleParttimeNotifyRequest) GetMessage() string {
    return r.message
}

