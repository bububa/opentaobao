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
type AliyunIndustryTttmProduceSyncAPIRequest struct {
    model.Params
    // 计划单
    _syncPlan   *SyncPlanDto
}

// 初始化AliyunIndustryTttmProduceSyncAPIRequest对象
func NewAliyunIndustryTttmProduceSyncRequest() *AliyunIndustryTttmProduceSyncAPIRequest{
    return &AliyunIndustryTttmProduceSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunIndustryTttmProduceSyncAPIRequest) GetApiMethodName() string {
    return "aliyun.industry.tttm.produce.sync"
}

// IRequest interface 方法, 获取API参数
func (r AliyunIndustryTttmProduceSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SyncPlan Setter
// 计划单
func (r *AliyunIndustryTttmProduceSyncAPIRequest) SetSyncPlan(_syncPlan *SyncPlanDto) error {
    r._syncPlan = _syncPlan
    r.Set("sync_plan", _syncPlan)
    return nil
}

// SyncPlan Getter
func (r AliyunIndustryTttmProduceSyncAPIRequest) GetSyncPlan() *SyncPlanDto {
    return r._syncPlan
}
