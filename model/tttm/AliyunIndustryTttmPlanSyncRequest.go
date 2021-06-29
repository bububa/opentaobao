package tttm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖生产计划单同步 API请求
aliyun.industry.tttm.plan.sync

ERP系统向天天特卖同步生产计划单的数据
*/
type AliyunIndustryTttmPlanSyncRequest struct {
    model.Params
    // 计划单
    syncPlan   *SyncPlanDto
}

// 初始化AliyunIndustryTttmPlanSyncRequest对象
func NewAliyunIndustryTttmPlanSyncRequest() *AliyunIndustryTttmPlanSyncRequest{
    return &AliyunIndustryTttmPlanSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunIndustryTttmPlanSyncRequest) GetApiMethodName() string {
    return "aliyun.industry.tttm.plan.sync"
}

// IRequest interface 方法, 获取API参数
func (r AliyunIndustryTttmPlanSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SyncPlan Setter
// 计划单
func (r *AliyunIndustryTttmPlanSyncRequest) SetSyncPlan(syncPlan *SyncPlanDto) error {
    r.syncPlan = syncPlan
    r.Set("sync_plan", syncPlan)
    return nil
}

// SyncPlan Getter
func (r AliyunIndustryTttmPlanSyncRequest) GetSyncPlan() *SyncPlanDto {
    return r.syncPlan
}
