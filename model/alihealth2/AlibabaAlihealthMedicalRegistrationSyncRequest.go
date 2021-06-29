package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康支付宝挂号记录回传接口 APIRequest
alibaba.alihealth.medical.registration.sync

阿里健康支付宝挂号记录回传接口
*/
type AlibabaAlihealthMedicalRegistrationSyncRequest struct {
    model.Params

    // 接口入参
    saveRequest   *CommonRequest4Top 

}

func NewAlibabaAlihealthMedicalRegistrationSyncRequest() *AlibabaAlihealthMedicalRegistrationSyncRequest{
    return &AlibabaAlihealthMedicalRegistrationSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMedicalRegistrationSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.registration.sync"
}

func (r AlibabaAlihealthMedicalRegistrationSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMedicalRegistrationSyncRequest) SetSaveRequest(saveRequest *CommonRequest4Top) error {
    r.saveRequest = saveRequest
    r.Set("save_request", saveRequest)
    return nil
}

func (r AlibabaAlihealthMedicalRegistrationSyncRequest) GetSaveRequest() *CommonRequest4Top {
    return r.saveRequest
}

