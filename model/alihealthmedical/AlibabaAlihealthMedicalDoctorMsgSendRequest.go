package alihealthmedical

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方医生消息写入 APIRequest
alibaba.alihealth.medical.doctor.msg.send

三方机构医生端发送消息同步写入阿里健康IM
*/
type AlibabaAlihealthMedicalDoctorMsgSendRequest struct {
    model.Params

    // request
    inquiry   *OuterMsgPullRequest 

}

func NewAlibabaAlihealthMedicalDoctorMsgSendRequest() *AlibabaAlihealthMedicalDoctorMsgSendRequest{
    return &AlibabaAlihealthMedicalDoctorMsgSendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMedicalDoctorMsgSendRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.doctor.msg.send"
}

func (r AlibabaAlihealthMedicalDoctorMsgSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMedicalDoctorMsgSendRequest) SetInquiry(inquiry *OuterMsgPullRequest) error {
    r.inquiry = inquiry
    r.Set("inquiry", inquiry)
    return nil
}

func (r AlibabaAlihealthMedicalDoctorMsgSendRequest) GetInquiry() *OuterMsgPullRequest {
    return r.inquiry
}

