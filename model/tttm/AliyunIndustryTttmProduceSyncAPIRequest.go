package tttm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunIndustryTttmProduceSyncAPIRequest 天天特卖生产进度同步 API请求
// aliyun.industry.tttm.produce.sync
//
// 天天特卖生产进度同步
type AliyunIndustryTttmProduceSyncAPIRequest struct {
	model.Params
	// 计划单
	_syncPlan *SyncPlanDto
}

// NewAliyunIndustryTttmProduceSyncRequest 初始化AliyunIndustryTttmProduceSyncAPIRequest对象
func NewAliyunIndustryTttmProduceSyncRequest() *AliyunIndustryTttmProduceSyncAPIRequest {
	return &AliyunIndustryTttmProduceSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliyunIndustryTttmProduceSyncAPIRequest) Reset() {
	r._syncPlan = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunIndustryTttmProduceSyncAPIRequest) GetApiMethodName() string {
	return "aliyun.industry.tttm.produce.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunIndustryTttmProduceSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunIndustryTttmProduceSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncPlan is SyncPlan Setter
// 计划单
func (r *AliyunIndustryTttmProduceSyncAPIRequest) SetSyncPlan(_syncPlan *SyncPlanDto) error {
	r._syncPlan = _syncPlan
	r.Set("sync_plan", _syncPlan)
	return nil
}

// GetSyncPlan SyncPlan Getter
func (r AliyunIndustryTttmProduceSyncAPIRequest) GetSyncPlan() *SyncPlanDto {
	return r._syncPlan
}

var poolAliyunIndustryTttmProduceSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAliyunIndustryTttmProduceSyncRequest()
	},
}

// GetAliyunIndustryTttmProduceSyncRequest 从 sync.Pool 获取 AliyunIndustryTttmProduceSyncAPIRequest
func GetAliyunIndustryTttmProduceSyncAPIRequest() *AliyunIndustryTttmProduceSyncAPIRequest {
	return poolAliyunIndustryTttmProduceSyncAPIRequest.Get().(*AliyunIndustryTttmProduceSyncAPIRequest)
}

// ReleaseAliyunIndustryTttmProduceSyncAPIRequest 将 AliyunIndustryTttmProduceSyncAPIRequest 放入 sync.Pool
func ReleaseAliyunIndustryTttmProduceSyncAPIRequest(v *AliyunIndustryTttmProduceSyncAPIRequest) {
	v.Reset()
	poolAliyunIndustryTttmProduceSyncAPIRequest.Put(v)
}
