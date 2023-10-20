package tttm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunIndustryTttmPlanSyncAPIRequest 天天特卖生产计划单同步 API请求
// aliyun.industry.tttm.plan.sync
//
// ERP系统向天天特卖同步生产计划单的数据
type AliyunIndustryTttmPlanSyncAPIRequest struct {
	model.Params
	// 计划单
	_syncPlan *SyncPlanDto
}

// NewAliyunIndustryTttmPlanSyncRequest 初始化AliyunIndustryTttmPlanSyncAPIRequest对象
func NewAliyunIndustryTttmPlanSyncRequest() *AliyunIndustryTttmPlanSyncAPIRequest {
	return &AliyunIndustryTttmPlanSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunIndustryTttmPlanSyncAPIRequest) GetApiMethodName() string {
	return "aliyun.industry.tttm.plan.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunIndustryTttmPlanSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunIndustryTttmPlanSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncPlan is SyncPlan Setter
// 计划单
func (r *AliyunIndustryTttmPlanSyncAPIRequest) SetSyncPlan(_syncPlan *SyncPlanDto) error {
	r._syncPlan = _syncPlan
	r.Set("sync_plan", _syncPlan)
	return nil
}

// GetSyncPlan SyncPlan Getter
func (r AliyunIndustryTttmPlanSyncAPIRequest) GetSyncPlan() *SyncPlanDto {
	return r._syncPlan
}
