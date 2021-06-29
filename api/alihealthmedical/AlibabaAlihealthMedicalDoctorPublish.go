package alihealthmedical

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthmedical"
)

/* 
三方机构医生信息上传 
alibaba.alihealth.medical.doctor.publish

三方机构上传医生信息
*/
func AlibabaAlihealthMedicalDoctorPublish(clt *core.SDKClient, req *alihealthmedical.AlibabaAlihealthMedicalDoctorPublishRequest, session string) (*alihealthmedical.AlibabaAlihealthMedicalDoctorPublishAPIResponse, error) {
    var resp alihealthmedical.AlibabaAlihealthMedicalDoctorPublishAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
