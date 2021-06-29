package idleparttime

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
兼职岗位同步 APIRequest
alibaba.idle.parttime.jobsync

服务商同步岗位信息给闲鱼
*/
type AlibabaIdleParttimeJobsyncRequest struct {
    model.Params

    // 岗位列表
    jobList   []PartTimeJob 

    // 同步数据的时间
    syncTime   int64 

}

func NewAlibabaIdleParttimeJobsyncRequest() *AlibabaIdleParttimeJobsyncRequest{
    return &AlibabaIdleParttimeJobsyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleParttimeJobsyncRequest) GetApiMethodName() string {
    return "alibaba.idle.parttime.jobsync"
}

func (r AlibabaIdleParttimeJobsyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleParttimeJobsyncRequest) SetJobList(jobList []PartTimeJob) error {
    r.jobList = jobList
    r.Set("job_list", jobList)
    return nil
}

func (r AlibabaIdleParttimeJobsyncRequest) GetJobList() []PartTimeJob {
    return r.jobList
}

func (r *AlibabaIdleParttimeJobsyncRequest) SetSyncTime(syncTime int64) error {
    r.syncTime = syncTime
    r.Set("sync_time", syncTime)
    return nil
}

func (r AlibabaIdleParttimeJobsyncRequest) GetSyncTime() int64 {
    return r.syncTime
}

