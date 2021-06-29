package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康预约挂号科室同步接口 APIRequest
alibaba.alihealth.medical.department.sync

阿里健康预约挂号科室同步接口
*/
type AlibabaAlihealthMedicalDepartmentSyncRequest struct {
    model.Params

    // 接口入参
    saveRequest   *CommonRequest4Top 

}

func NewAlibabaAlihealthMedicalDepartmentSyncRequest() *AlibabaAlihealthMedicalDepartmentSyncRequest{
    return &AlibabaAlihealthMedicalDepartmentSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMedicalDepartmentSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.department.sync"
}

func (r AlibabaAlihealthMedicalDepartmentSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMedicalDepartmentSyncRequest) SetSaveRequest(saveRequest *CommonRequest4Top) error {
    r.saveRequest = saveRequest
    r.Set("save_request", saveRequest)
    return nil
}

func (r AlibabaAlihealthMedicalDepartmentSyncRequest) GetSaveRequest() *CommonRequest4Top {
    return r.saveRequest
}

