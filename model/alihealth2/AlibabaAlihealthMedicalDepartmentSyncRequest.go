package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康预约挂号科室同步接口 API请求
alibaba.alihealth.medical.department.sync

阿里健康预约挂号科室同步接口
*/
type AlibabaAlihealthMedicalDepartmentSyncRequest struct {
    model.Params
    // 接口入参
    saveRequest   *CommonRequest4Top
}

// 初始化AlibabaAlihealthMedicalDepartmentSyncRequest对象
func NewAlibabaAlihealthMedicalDepartmentSyncRequest() *AlibabaAlihealthMedicalDepartmentSyncRequest{
    return &AlibabaAlihealthMedicalDepartmentSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalDepartmentSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.department.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalDepartmentSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SaveRequest Setter
// 接口入参
func (r *AlibabaAlihealthMedicalDepartmentSyncRequest) SetSaveRequest(saveRequest *CommonRequest4Top) error {
    r.saveRequest = saveRequest
    r.Set("save_request", saveRequest)
    return nil
}

// SaveRequest Getter
func (r AlibabaAlihealthMedicalDepartmentSyncRequest) GetSaveRequest() *CommonRequest4Top {
    return r.saveRequest
}
