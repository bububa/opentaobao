package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
微医数据号源回传 API请求
alibaba.alihealth.medical.register.weiyi.sync

微医号源数据回传
*/
type AlibabaAlihealthMedicalRegisterWeiyiSyncRequest struct {
    model.Params
    // 号源数据实体
    _serviceRequest   *SourcesReturnVo
}

// 初始化AlibabaAlihealthMedicalRegisterWeiyiSyncRequest对象
func NewAlibabaAlihealthMedicalRegisterWeiyiSyncRequest() *AlibabaAlihealthMedicalRegisterWeiyiSyncRequest{
    return &AlibabaAlihealthMedicalRegisterWeiyiSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalRegisterWeiyiSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.register.weiyi.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalRegisterWeiyiSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceRequest Setter
// 号源数据实体
func (r *AlibabaAlihealthMedicalRegisterWeiyiSyncRequest) SetServiceRequest(_serviceRequest *SourcesReturnVo) error {
    r._serviceRequest = _serviceRequest
    r.Set("service_request", _serviceRequest)
    return nil
}

// ServiceRequest Getter
func (r AlibabaAlihealthMedicalRegisterWeiyiSyncRequest) GetServiceRequest() *SourcesReturnVo {
    return r._serviceRequest
}
