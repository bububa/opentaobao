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
type AlibabaAlihealthOutflowVisitinfoSyncAPIRequest struct {
    model.Params
    // 入参
    _syncVisitInfoRequest   *SyncVisitInfoRequest
}

// 初始化AlibabaAlihealthOutflowVisitinfoSyncAPIRequest对象
func NewAlibabaAlihealthOutflowVisitinfoSyncRequest() *AlibabaAlihealthOutflowVisitinfoSyncAPIRequest{
    return &AlibabaAlihealthOutflowVisitinfoSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowVisitinfoSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.visitinfo.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowVisitinfoSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SyncVisitInfoRequest Setter
// 入参
func (r *AlibabaAlihealthOutflowVisitinfoSyncAPIRequest) SetSyncVisitInfoRequest(_syncVisitInfoRequest *SyncVisitInfoRequest) error {
    r._syncVisitInfoRequest = _syncVisitInfoRequest
    r.Set("sync_visit_info_request", _syncVisitInfoRequest)
    return nil
}

// SyncVisitInfoRequest Getter
func (r AlibabaAlihealthOutflowVisitinfoSyncAPIRequest) GetSyncVisitInfoRequest() *SyncVisitInfoRequest {
    return r._syncVisitInfoRequest
}
