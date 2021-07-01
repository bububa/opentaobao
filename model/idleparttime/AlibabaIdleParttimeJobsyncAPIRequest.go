package idleparttime

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleParttimeJobsyncAPIRequest
兼职岗位同步 API请求
alibaba.idle.parttime.jobsync

服务商同步岗位信息给闲鱼 */
type AlibabaIdleParttimeJobsyncAPIRequest struct {
	model.Params
	// 岗位列表
	_jobList []PartTimeJob
	// 同步数据的时间
	_syncTime int64
}

// NewAlibabaIdleParttimeJobsyncRequest 初始化AlibabaIdleParttimeJobsyncAPIRequest对象
func NewAlibabaIdleParttimeJobsyncRequest() *AlibabaIdleParttimeJobsyncAPIRequest {
	return &AlibabaIdleParttimeJobsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleParttimeJobsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.parttime.jobsync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleParttimeJobsyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is JobList Setter
// 岗位列表
func (r *AlibabaIdleParttimeJobsyncAPIRequest) SetJobList(_jobList []PartTimeJob) error {
	r._jobList = _jobList
	r.Set("job_list", _jobList)
	return nil
}

// Get JobList Getter
func (r AlibabaIdleParttimeJobsyncAPIRequest) GetJobList() []PartTimeJob {
	return r._jobList
}

// Set is SyncTime Setter
// 同步数据的时间
func (r *AlibabaIdleParttimeJobsyncAPIRequest) SetSyncTime(_syncTime int64) error {
	r._syncTime = _syncTime
	r.Set("sync_time", _syncTime)
	return nil
}

// Get SyncTime Getter
func (r AlibabaIdleParttimeJobsyncAPIRequest) GetSyncTime() int64 {
	return r._syncTime
}
