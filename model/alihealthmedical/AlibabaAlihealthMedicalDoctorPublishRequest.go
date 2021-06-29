package alihealthmedical

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方机构医生信息上传 APIRequest
alibaba.alihealth.medical.doctor.publish

三方机构上传医生信息
*/
type AlibabaAlihealthMedicalDoctorPublishRequest struct {
    model.Params

    // 三方机构医生上传request
    outerDoctorPublishRequest   *OuterDoctorPublishRequest 

}

func NewAlibabaAlihealthMedicalDoctorPublishRequest() *AlibabaAlihealthMedicalDoctorPublishRequest{
    return &AlibabaAlihealthMedicalDoctorPublishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMedicalDoctorPublishRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.doctor.publish"
}

func (r AlibabaAlihealthMedicalDoctorPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMedicalDoctorPublishRequest) SetOuterDoctorPublishRequest(outerDoctorPublishRequest *OuterDoctorPublishRequest) error {
    r.outerDoctorPublishRequest = outerDoctorPublishRequest
    r.Set("outer_doctor_publish_request", outerDoctorPublishRequest)
    return nil
}

func (r AlibabaAlihealthMedicalDoctorPublishRequest) GetOuterDoctorPublishRequest() *OuterDoctorPublishRequest {
    return r.outerDoctorPublishRequest
}

