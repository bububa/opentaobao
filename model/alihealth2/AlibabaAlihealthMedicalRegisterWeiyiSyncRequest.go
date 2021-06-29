package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
微医数据号源回传 APIRequest
alibaba.alihealth.medical.register.weiyi.sync

微医号源数据回传
*/
type AlibabaAlihealthMedicalRegisterWeiyiSyncRequest struct {
    model.Params

    // 号源数据实体
    serviceRequest   *SourcesReturnVo 

}

func NewAlibabaAlihealthMedicalRegisterWeiyiSyncRequest() *AlibabaAlihealthMedicalRegisterWeiyiSyncRequest{
    return &AlibabaAlihealthMedicalRegisterWeiyiSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMedicalRegisterWeiyiSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.register.weiyi.sync"
}

func (r AlibabaAlihealthMedicalRegisterWeiyiSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMedicalRegisterWeiyiSyncRequest) SetServiceRequest(serviceRequest *SourcesReturnVo) error {
    r.serviceRequest = serviceRequest
    r.Set("service_request", serviceRequest)
    return nil
}

func (r AlibabaAlihealthMedicalRegisterWeiyiSyncRequest) GetServiceRequest() *SourcesReturnVo {
    return r.serviceRequest
}

