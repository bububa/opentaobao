package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里将康支付宝挂号数据医院回传接口 APIRequest
alibaba.alihealth.medical.hospital.sync

阿里将康支付宝挂号数据医院回传接口
*/
type AlibabaAlihealthMedicalHospitalSyncRequest struct {
    model.Params

    // top保存入参
    saveRequest   *CommonRequest4Top 

}

func NewAlibabaAlihealthMedicalHospitalSyncRequest() *AlibabaAlihealthMedicalHospitalSyncRequest{
    return &AlibabaAlihealthMedicalHospitalSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMedicalHospitalSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.hospital.sync"
}

func (r AlibabaAlihealthMedicalHospitalSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMedicalHospitalSyncRequest) SetSaveRequest(saveRequest *CommonRequest4Top) error {
    r.saveRequest = saveRequest
    r.Set("save_request", saveRequest)
    return nil
}

func (r AlibabaAlihealthMedicalHospitalSyncRequest) GetSaveRequest() *CommonRequest4Top {
    return r.saveRequest
}

