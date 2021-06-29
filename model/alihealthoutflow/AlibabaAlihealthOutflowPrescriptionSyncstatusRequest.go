package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-处方状态同步 APIRequest
alibaba.alihealth.outflow.prescription.syncstatus

阿里健康-处方外流-对外提供同步处方状态功能
*/
type AlibabaAlihealthOutflowPrescriptionSyncstatusRequest struct {
    model.Params

    // 入参
    syncStatusRequest   *SyncPrescriptionStatusRequest 

}

func NewAlibabaAlihealthOutflowPrescriptionSyncstatusRequest() *AlibabaAlihealthOutflowPrescriptionSyncstatusRequest{
    return &AlibabaAlihealthOutflowPrescriptionSyncstatusRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthOutflowPrescriptionSyncstatusRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.prescription.syncstatus"
}

func (r AlibabaAlihealthOutflowPrescriptionSyncstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthOutflowPrescriptionSyncstatusRequest) SetSyncStatusRequest(syncStatusRequest *SyncPrescriptionStatusRequest) error {
    r.syncStatusRequest = syncStatusRequest
    r.Set("sync_status_request", syncStatusRequest)
    return nil
}

func (r AlibabaAlihealthOutflowPrescriptionSyncstatusRequest) GetSyncStatusRequest() *SyncPrescriptionStatusRequest {
    return r.syncStatusRequest
}

