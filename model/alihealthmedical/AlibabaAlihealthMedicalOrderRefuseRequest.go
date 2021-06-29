package alihealthmedical

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方机构通知平台"医生拒诊" APIRequest
alibaba.alihealth.medical.order.refuse

三方机构通知平台"医生拒诊"
*/
type AlibabaAlihealthMedicalOrderRefuseRequest struct {
    model.Params

    // 请求入参
    requestInfo   *RefuseOrderRequestDTO 

}

func NewAlibabaAlihealthMedicalOrderRefuseRequest() *AlibabaAlihealthMedicalOrderRefuseRequest{
    return &AlibabaAlihealthMedicalOrderRefuseRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMedicalOrderRefuseRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.order.refuse"
}

func (r AlibabaAlihealthMedicalOrderRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMedicalOrderRefuseRequest) SetRequestInfo(requestInfo *RefuseOrderRequestDTO) error {
    r.requestInfo = requestInfo
    r.Set("request_info", requestInfo)
    return nil
}

func (r AlibabaAlihealthMedicalOrderRefuseRequest) GetRequestInfo() *RefuseOrderRequestDTO {
    return r.requestInfo
}

