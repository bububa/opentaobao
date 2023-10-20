package tttm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunindustrytttmproducesyncAPIRequest 天天特卖生产进度同步 API请求
// aliyun.industry.tttm.produce.sync
//
// 天天特卖生产进度同步
type AliyunindustrytttmproducesyncAPIRequest struct {
	model.Params
	// 计划单
	_syncPlan *SyncPlanDto
}

// NewAliyunindustrytttmproducesyncRequest 初始化AliyunindustrytttmproducesyncAPIRequest对象
func NewAliyunindustrytttmproducesyncRequest() *AliyunindustrytttmproducesyncAPIRequest {
	return &AliyunindustrytttmproducesyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunindustrytttmproducesyncAPIRequest) GetApiMethodName() string {
	return "aliyun.industry.tttm.produce.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunindustrytttmproducesyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunindustrytttmproducesyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncPlan is SyncPlan Setter
// 计划单
func (r *AliyunindustrytttmproducesyncAPIRequest) SetSyncPlan(_syncPlan *SyncPlanDto) error {
	r._syncPlan = _syncPlan
	r.Set("sync_plan", _syncPlan)
	return nil
}

// GetSyncPlan SyncPlan Getter
func (r AliyunindustrytttmproducesyncAPIRequest) GetSyncPlan() *SyncPlanDto {
	return r._syncPlan
}
