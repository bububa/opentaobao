package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-问诊、处方同步 API请求
alibaba.alihealth.outflow.visitinfo.sync

阿里健康-处方外流-对外提供问诊、处方功能
*/
type AlibabaAlihealthOutflowVisitinfoSyncRequest struct {
    model.Params
    // 入参
    syncVisitInfoRequest   *SyncVisitInfoRequest
}

// 初始化AlibabaAlihealthOutflowVisitinfoSyncRequest对象
func NewAlibabaAlihealthOutflowVisitinfoSyncRequest() *AlibabaAlihealthOutflowVisitinfoSyncRequest{
    return &AlibabaAlihealthOutflowVisitinfoSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowVisitinfoSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.visitinfo.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowVisitinfoSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SyncVisitInfoRequest Setter
// 入参
func (r *AlibabaAlihealthOutflowVisitinfoSyncRequest) SetSyncVisitInfoRequest(syncVisitInfoRequest *SyncVisitInfoRequest) error {
    r.syncVisitInfoRequest = syncVisitInfoRequest
    r.Set("sync_visit_info_request", syncVisitInfoRequest)
    return nil
}

// SyncVisitInfoRequest Getter
func (r AlibabaAlihealthOutflowVisitinfoSyncRequest) GetSyncVisitInfoRequest() *SyncVisitInfoRequest {
    return r.syncVisitInfoRequest
}
