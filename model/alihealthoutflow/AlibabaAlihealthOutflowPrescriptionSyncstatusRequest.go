package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-处方状态同步 API请求
alibaba.alihealth.outflow.prescription.syncstatus

阿里健康-处方外流-对外提供同步处方状态功能
*/
type AlibabaAlihealthOutflowPrescriptionSyncstatusRequest struct {
    model.Params
    // 入参
    _syncStatusRequest   *SyncPrescriptionStatusRequest
}

// 初始化AlibabaAlihealthOutflowPrescriptionSyncstatusRequest对象
func NewAlibabaAlihealthOutflowPrescriptionSyncstatusRequest() *AlibabaAlihealthOutflowPrescriptionSyncstatusRequest{
    return &AlibabaAlihealthOutflowPrescriptionSyncstatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowPrescriptionSyncstatusRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.prescription.syncstatus"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowPrescriptionSyncstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SyncStatusRequest Setter
// 入参
func (r *AlibabaAlihealthOutflowPrescriptionSyncstatusRequest) SetSyncStatusRequest(_syncStatusRequest *SyncPrescriptionStatusRequest) error {
    r._syncStatusRequest = _syncStatusRequest
    r.Set("sync_status_request", _syncStatusRequest)
    return nil
}

// SyncStatusRequest Getter
func (r AlibabaAlihealthOutflowPrescriptionSyncstatusRequest) GetSyncStatusRequest() *SyncPrescriptionStatusRequest {
    return r._syncStatusRequest
}
