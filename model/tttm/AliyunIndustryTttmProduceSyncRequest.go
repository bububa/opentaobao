package tttm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖生产进度同步 APIRequest
aliyun.industry.tttm.produce.sync

天天特卖生产进度同步
*/
type AliyunIndustryTttmProduceSyncRequest struct {
    model.Params

    // 计划单
    syncPlan   *SyncPlanDto 

}

func NewAliyunIndustryTttmProduceSyncRequest() *AliyunIndustryTttmProduceSyncRequest{
    return &AliyunIndustryTttmProduceSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AliyunIndustryTttmProduceSyncRequest) GetApiMethodName() string {
    return "aliyun.industry.tttm.produce.sync"
}

func (r AliyunIndustryTttmProduceSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliyunIndustryTttmProduceSyncRequest) SetSyncPlan(syncPlan *SyncPlanDto) error {
    r.syncPlan = syncPlan
    r.Set("sync_plan", syncPlan)
    return nil
}

func (r AliyunIndustryTttmProduceSyncRequest) GetSyncPlan() *SyncPlanDto {
    return r.syncPlan
}

