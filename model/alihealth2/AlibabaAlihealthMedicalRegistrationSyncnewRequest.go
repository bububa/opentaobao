package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康新挂号数据回传 APIRequest
alibaba.alihealth.medical.registration.syncnew

阿里健康新挂号记录回传接口
*/
type AlibabaAlihealthMedicalRegistrationSyncnewRequest struct {
    model.Params

    // 接口入参
    saveRequest   *CommonRequest4Top 

}

func NewAlibabaAlihealthMedicalRegistrationSyncnewRequest() *AlibabaAlihealthMedicalRegistrationSyncnewRequest{
    return &AlibabaAlihealthMedicalRegistrationSyncnewRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMedicalRegistrationSyncnewRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.registration.syncnew"
}

func (r AlibabaAlihealthMedicalRegistrationSyncnewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMedicalRegistrationSyncnewRequest) SetSaveRequest(saveRequest *CommonRequest4Top) error {
    r.saveRequest = saveRequest
    r.Set("save_request", saveRequest)
    return nil
}

func (r AlibabaAlihealthMedicalRegistrationSyncnewRequest) GetSaveRequest() *CommonRequest4Top {
    return r.saveRequest
}

