package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康支付宝挂号记录回传接口 API请求
alibaba.alihealth.medical.registration.sync

阿里健康支付宝挂号记录回传接口
*/
type AlibabaAlihealthMedicalRegistrationSyncRequest struct {
    model.Params
    // 接口入参
    _saveRequest   *CommonRequest4Top
}

// 初始化AlibabaAlihealthMedicalRegistrationSyncRequest对象
func NewAlibabaAlihealthMedicalRegistrationSyncRequest() *AlibabaAlihealthMedicalRegistrationSyncRequest{
    return &AlibabaAlihealthMedicalRegistrationSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalRegistrationSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.registration.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalRegistrationSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SaveRequest Setter
// 接口入参
func (r *AlibabaAlihealthMedicalRegistrationSyncRequest) SetSaveRequest(_saveRequest *CommonRequest4Top) error {
    r._saveRequest = _saveRequest
    r.Set("save_request", _saveRequest)
    return nil
}

// SaveRequest Getter
func (r AlibabaAlihealthMedicalRegistrationSyncRequest) GetSaveRequest() *CommonRequest4Top {
    return r._saveRequest
}
