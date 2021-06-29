package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-患者基础信息同步 APIRequest
alibaba.alihealth.outflow.patientinfo.sync

阿里健康-处方外流-对外提供同步患者基础信息功能
*/
type AlibabaAlihealthOutflowPatientinfoSyncRequest struct {
    model.Params

    // 入参
    syncPatientInfoRequest   *SyncPatientInfoRequest 

}

func NewAlibabaAlihealthOutflowPatientinfoSyncRequest() *AlibabaAlihealthOutflowPatientinfoSyncRequest{
    return &AlibabaAlihealthOutflowPatientinfoSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthOutflowPatientinfoSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.patientinfo.sync"
}

func (r AlibabaAlihealthOutflowPatientinfoSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthOutflowPatientinfoSyncRequest) SetSyncPatientInfoRequest(syncPatientInfoRequest *SyncPatientInfoRequest) error {
    r.syncPatientInfoRequest = syncPatientInfoRequest
    r.Set("sync_patient_info_request", syncPatientInfoRequest)
    return nil
}

func (r AlibabaAlihealthOutflowPatientinfoSyncRequest) GetSyncPatientInfoRequest() *SyncPatientInfoRequest {
    return r.syncPatientInfoRequest
}

