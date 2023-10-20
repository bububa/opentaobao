package tttm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunindustrytttmplansyncAPIRequest 天天特卖生产计划单同步 API请求
// aliyun.industry.tttm.plan.sync
//
// ERP系统向天天特卖同步生产计划单的数据
type AliyunindustrytttmplansyncAPIRequest struct {
	model.Params
	// 计划单
	_syncPlan *SyncPlanDto
}

// NewAliyunindustrytttmplansyncRequest 初始化AliyunindustrytttmplansyncAPIRequest对象
func NewAliyunindustrytttmplansyncRequest() *AliyunindustrytttmplansyncAPIRequest {
	return &AliyunindustrytttmplansyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunindustrytttmplansyncAPIRequest) GetApiMethodName() string {
	return "aliyun.industry.tttm.plan.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunindustrytttmplansyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunindustrytttmplansyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncPlan is SyncPlan Setter
// 计划单
func (r *AliyunindustrytttmplansyncAPIRequest) SetSyncPlan(_syncPlan *SyncPlanDto) error {
	r._syncPlan = _syncPlan
	r.Set("sync_plan", _syncPlan)
	return nil
}

// GetSyncPlan SyncPlan Getter
func (r AliyunindustrytttmplansyncAPIRequest) GetSyncPlan() *SyncPlanDto {
	return r._syncPlan
}
