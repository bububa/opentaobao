package alihealthmedical

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方机构医生信息上传 API请求
alibaba.alihealth.medical.doctor.publish

三方机构上传医生信息
*/
type AlibabaAlihealthMedicalDoctorPublishRequest struct {
    model.Params
    // 三方机构医生上传request
    outerDoctorPublishRequest   *OuterDoctorPublishRequest
}

// 初始化AlibabaAlihealthMedicalDoctorPublishRequest对象
func NewAlibabaAlihealthMedicalDoctorPublishRequest() *AlibabaAlihealthMedicalDoctorPublishRequest{
    return &AlibabaAlihealthMedicalDoctorPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalDoctorPublishRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.doctor.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalDoctorPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterDoctorPublishRequest Setter
// 三方机构医生上传request
func (r *AlibabaAlihealthMedicalDoctorPublishRequest) SetOuterDoctorPublishRequest(outerDoctorPublishRequest *OuterDoctorPublishRequest) error {
    r.outerDoctorPublishRequest = outerDoctorPublishRequest
    r.Set("outer_doctor_publish_request", outerDoctorPublishRequest)
    return nil
}

// OuterDoctorPublishRequest Getter
func (r AlibabaAlihealthMedicalDoctorPublishRequest) GetOuterDoctorPublishRequest() *OuterDoctorPublishRequest {
    return r.outerDoctorPublishRequest
}
