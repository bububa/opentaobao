package alihealth

import (
    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康处方平台获取授权码 APIResponse
alibaba.alihealth.prescription.auth.get

获取处方用户授权
*/
type AlibabaAlihealthPrescriptionAuthGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlihealthPrescriptionAuthGetResponse `json:"alibaba_alihealth_prescription_auth_get_response,omitempty"`
}

type AlibabaAlihealthPrescriptionAuthGetResponse struct {

    // alinkappserver系统返回的通用结果类
    ServiceResult   *ServiceResult `json:"service_result,omitempty"`

}
