package idleparttime

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
兼职岗位同步 API请求
alibaba.idle.parttime.jobsync

服务商同步岗位信息给闲鱼
*/
type AlibabaIdleParttimeJobsyncRequest struct {
    model.Params
    // 岗位列表
    _jobList   []PartTimeJob
    // 同步数据的时间
    _syncTime   int64
}

// 初始化AlibabaIdleParttimeJobsyncRequest对象
func NewAlibabaIdleParttimeJobsyncRequest() *AlibabaIdleParttimeJobsyncRequest{
    return &AlibabaIdleParttimeJobsyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleParttimeJobsyncRequest) GetApiMethodName() string {
    return "alibaba.idle.parttime.jobsync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleParttimeJobsyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// JobList Setter
// 岗位列表
func (r *AlibabaIdleParttimeJobsyncRequest) SetJobList(_jobList []PartTimeJob) error {
    r._jobList = _jobList
    r.Set("job_list", _jobList)
    return nil
}

// JobList Getter
func (r AlibabaIdleParttimeJobsyncRequest) GetJobList() []PartTimeJob {
    return r._jobList
}
// SyncTime Setter
// 同步数据的时间
func (r *AlibabaIdleParttimeJobsyncRequest) SetSyncTime(_syncTime int64) error {
    r._syncTime = _syncTime
    r.Set("sync_time", _syncTime)
    return nil
}

// SyncTime Getter
func (r AlibabaIdleParttimeJobsyncRequest) GetSyncTime() int64 {
    return r._syncTime
}
