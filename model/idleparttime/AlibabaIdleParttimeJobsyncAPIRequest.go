package idleparttime

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleParttimeJobsyncAPIRequest 兼职岗位同步 API请求
// alibaba.idle.parttime.jobsync
//
// 服务商同步岗位信息给闲鱼
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleParttimeJobsyncAPIRequest) Reset() {
	r._jobList = r._jobList[:0]
	r._syncTime = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleParttimeJobsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.parttime.jobsync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleParttimeJobsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleParttimeJobsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetJobList is JobList Setter
// 岗位列表
func (r *AlibabaIdleParttimeJobsyncAPIRequest) SetJobList(_jobList []PartTimeJob) error {
	r._jobList = _jobList
	r.Set("job_list", _jobList)
	return nil
}

// GetJobList JobList Getter
func (r AlibabaIdleParttimeJobsyncAPIRequest) GetJobList() []PartTimeJob {
	return r._jobList
}

// SetSyncTime is SyncTime Setter
// 同步数据的时间
func (r *AlibabaIdleParttimeJobsyncAPIRequest) SetSyncTime(_syncTime int64) error {
	r._syncTime = _syncTime
	r.Set("sync_time", _syncTime)
	return nil
}

// GetSyncTime SyncTime Getter
func (r AlibabaIdleParttimeJobsyncAPIRequest) GetSyncTime() int64 {
	return r._syncTime
}

var poolAlibabaIdleParttimeJobsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleParttimeJobsyncRequest()
	},
}

// GetAlibabaIdleParttimeJobsyncRequest 从 sync.Pool 获取 AlibabaIdleParttimeJobsyncAPIRequest
func GetAlibabaIdleParttimeJobsyncAPIRequest() *AlibabaIdleParttimeJobsyncAPIRequest {
	return poolAlibabaIdleParttimeJobsyncAPIRequest.Get().(*AlibabaIdleParttimeJobsyncAPIRequest)
}

// ReleaseAlibabaIdleParttimeJobsyncAPIRequest 将 AlibabaIdleParttimeJobsyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleParttimeJobsyncAPIRequest(v *AlibabaIdleParttimeJobsyncAPIRequest) {
	v.Reset()
	poolAlibabaIdleParttimeJobsyncAPIRequest.Put(v)
}
