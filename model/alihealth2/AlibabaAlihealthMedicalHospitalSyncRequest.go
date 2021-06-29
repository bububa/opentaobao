package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里将康支付宝挂号数据医院回传接口 API请求
alibaba.alihealth.medical.hospital.sync

阿里将康支付宝挂号数据医院回传接口
*/
type AlibabaAlihealthMedicalHospitalSyncRequest struct {
    model.Params
    // top保存入参
    saveRequest   *CommonRequest4Top
}

// 初始化AlibabaAlihealthMedicalHospitalSyncRequest对象
func NewAlibabaAlihealthMedicalHospitalSyncRequest() *AlibabaAlihealthMedicalHospitalSyncRequest{
    return &AlibabaAlihealthMedicalHospitalSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalHospitalSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.hospital.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalHospitalSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SaveRequest Setter
// top保存入参
func (r *AlibabaAlihealthMedicalHospitalSyncRequest) SetSaveRequest(saveRequest *CommonRequest4Top) error {
    r.saveRequest = saveRequest
    r.Set("save_request", saveRequest)
    return nil
}

// SaveRequest Getter
func (r AlibabaAlihealthMedicalHospitalSyncRequest) GetSaveRequest() *CommonRequest4Top {
    return r.saveRequest
}
