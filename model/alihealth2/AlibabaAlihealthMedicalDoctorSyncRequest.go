package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康预约挂号医生同步接口 APIRequest
alibaba.alihealth.medical.doctor.sync

阿里健康预约挂号医生同步接口
*/
type AlibabaAlihealthMedicalDoctorSyncRequest struct {
    model.Params

    // 接口入参
    saveRequest   *CommonRequest4Top 

}

func NewAlibabaAlihealthMedicalDoctorSyncRequest() *AlibabaAlihealthMedicalDoctorSyncRequest{
    return &AlibabaAlihealthMedicalDoctorSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMedicalDoctorSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.doctor.sync"
}

func (r AlibabaAlihealthMedicalDoctorSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMedicalDoctorSyncRequest) SetSaveRequest(saveRequest *CommonRequest4Top) error {
    r.saveRequest = saveRequest
    r.Set("save_request", saveRequest)
    return nil
}

func (r AlibabaAlihealthMedicalDoctorSyncRequest) GetSaveRequest() *CommonRequest4Top {
    return r.saveRequest
}

