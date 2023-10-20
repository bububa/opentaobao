package idleparttime

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleparttimejobsyncAPIRequest 兼职岗位同步 API请求
// alibaba.idle.parttime.jobsync
//
// 服务商同步岗位信息给闲鱼
type AlibabaidleparttimejobsyncAPIRequest struct {
	model.Params
	// 岗位列表
	_jobList []PartTimeJob
	// 同步数据的时间
	_syncTime int64
}

// NewAlibabaidleparttimejobsyncRequest 初始化AlibabaidleparttimejobsyncAPIRequest对象
func NewAlibabaidleparttimejobsyncRequest() *AlibabaidleparttimejobsyncAPIRequest {
	return &AlibabaidleparttimejobsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleparttimejobsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.parttime.jobsync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleparttimejobsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleparttimejobsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetJobList is JobList Setter
// 岗位列表
func (r *AlibabaidleparttimejobsyncAPIRequest) SetJobList(_jobList []PartTimeJob) error {
	r._jobList = _jobList
	r.Set("job_list", _jobList)
	return nil
}

// GetJobList JobList Getter
func (r AlibabaidleparttimejobsyncAPIRequest) GetJobList() []PartTimeJob {
	return r._jobList
}

// SetSyncTime is SyncTime Setter
// 同步数据的时间
func (r *AlibabaidleparttimejobsyncAPIRequest) SetSyncTime(_syncTime int64) error {
	r._syncTime = _syncTime
	r.Set("sync_time", _syncTime)
	return nil
}

// GetSyncTime SyncTime Getter
func (r AlibabaidleparttimejobsyncAPIRequest) GetSyncTime() int64 {
	return r._syncTime
}
