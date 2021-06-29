package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-问诊、处方同步 APIRequest
alibaba.alihealth.outflow.visitinfo.sync

阿里健康-处方外流-对外提供问诊、处方功能
*/
type AlibabaAlihealthOutflowVisitinfoSyncRequest struct {
    model.Params

    // 入参
    syncVisitInfoRequest   *SyncVisitInfoRequest 

}

func NewAlibabaAlihealthOutflowVisitinfoSyncRequest() *AlibabaAlihealthOutflowVisitinfoSyncRequest{
    return &AlibabaAlihealthOutflowVisitinfoSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthOutflowVisitinfoSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.visitinfo.sync"
}

func (r AlibabaAlihealthOutflowVisitinfoSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthOutflowVisitinfoSyncRequest) SetSyncVisitInfoRequest(syncVisitInfoRequest *SyncVisitInfoRequest) error {
    r.syncVisitInfoRequest = syncVisitInfoRequest
    r.Set("sync_visit_info_request", syncVisitInfoRequest)
    return nil
}

func (r AlibabaAlihealthOutflowVisitinfoSyncRequest) GetSyncVisitInfoRequest() *SyncVisitInfoRequest {
    return r.syncVisitInfoRequest
}

