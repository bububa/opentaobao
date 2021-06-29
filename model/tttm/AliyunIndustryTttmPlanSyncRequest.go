package tttm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖生产计划单同步 APIRequest
aliyun.industry.tttm.plan.sync

ERP系统向天天特卖同步生产计划单的数据
*/
type AliyunIndustryTttmPlanSyncRequest struct {
    model.Params

    // 计划单
    syncPlan   *SyncPlanDto 

}

func NewAliyunIndustryTttmPlanSyncRequest() *AliyunIndustryTttmPlanSyncRequest{
    return &AliyunIndustryTttmPlanSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AliyunIndustryTttmPlanSyncRequest) GetApiMethodName() string {
    return "aliyun.industry.tttm.plan.sync"
}

func (r AliyunIndustryTttmPlanSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliyunIndustryTttmPlanSyncRequest) SetSyncPlan(syncPlan *SyncPlanDto) error {
    r.syncPlan = syncPlan
    r.Set("sync_plan", syncPlan)
    return nil
}

func (r AliyunIndustryTttmPlanSyncRequest) GetSyncPlan() *SyncPlanDto {
    return r.syncPlan
}

