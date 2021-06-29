package alihealthmedical

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方医生消息写入 API请求
alibaba.alihealth.medical.doctor.msg.send

三方机构医生端发送消息同步写入阿里健康IM
*/
type AlibabaAlihealthMedicalDoctorMsgSendRequest struct {
    model.Params
    // request
    inquiry   *OuterMsgPullRequest
}

// 初始化AlibabaAlihealthMedicalDoctorMsgSendRequest对象
func NewAlibabaAlihealthMedicalDoctorMsgSendRequest() *AlibabaAlihealthMedicalDoctorMsgSendRequest{
    return &AlibabaAlihealthMedicalDoctorMsgSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalDoctorMsgSendRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.doctor.msg.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalDoctorMsgSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Inquiry Setter
// request
func (r *AlibabaAlihealthMedicalDoctorMsgSendRequest) SetInquiry(inquiry *OuterMsgPullRequest) error {
    r.inquiry = inquiry
    r.Set("inquiry", inquiry)
    return nil
}

// Inquiry Getter
func (r AlibabaAlihealthMedicalDoctorMsgSendRequest) GetInquiry() *OuterMsgPullRequest {
    return r.inquiry
}
