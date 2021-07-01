package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康预约挂号医生同步接口 API请求
alibaba.alihealth.medical.doctor.sync

阿里健康预约挂号医生同步接口
*/
type AlibabaAlihealthMedicalDoctorSyncAPIRequest struct {
    model.Params
    // 接口入参
    _saveRequest   *CommonRequest4Top
}

// 初始化AlibabaAlihealthMedicalDoctorSyncAPIRequest对象
func NewAlibabaAlihealthMedicalDoctorSyncRequest() *AlibabaAlihealthMedicalDoctorSyncAPIRequest{
    return &AlibabaAlihealthMedicalDoctorSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalDoctorSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.doctor.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalDoctorSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SaveRequest Setter
// 接口入参
func (r *AlibabaAlihealthMedicalDoctorSyncAPIRequest) SetSaveRequest(_saveRequest *CommonRequest4Top) error {
    r._saveRequest = _saveRequest
    r.Set("save_request", _saveRequest)
    return nil
}

// SaveRequest Getter
func (r AlibabaAlihealthMedicalDoctorSyncAPIRequest) GetSaveRequest() *CommonRequest4Top {
    return r._saveRequest
}
