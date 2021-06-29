package tttm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖生产进度同步 API请求
aliyun.industry.tttm.produce.sync

天天特卖生产进度同步
*/
type AliyunIndustryTttmProduceSyncRequest struct {
    model.Params
    // 计划单
    _syncPlan   *SyncPlanDto
}

// 初始化AliyunIndustryTttmProduceSyncRequest对象
func NewAliyunIndustryTttmProduceSyncRequest() *AliyunIndustryTttmProduceSyncRequest{
    return &AliyunIndustryTttmProduceSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunIndustryTttmProduceSyncRequest) GetApiMethodName() string {
    return "aliyun.industry.tttm.produce.sync"
}

// IRequest interface 方法, 获取API参数
func (r AliyunIndustryTttmProduceSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SyncPlan Setter
// 计划单
func (r *AliyunIndustryTttmProduceSyncRequest) SetSyncPlan(_syncPlan *SyncPlanDto) error {
    r._syncPlan = _syncPlan
    r.Set("sync_plan", _syncPlan)
    return nil
}

// SyncPlan Getter
func (r AliyunIndustryTttmProduceSyncRequest) GetSyncPlan() *SyncPlanDto {
    return r._syncPlan
}
