package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
保护结果反馈 APIRequest
alibaba.security.jaq.spamregisterprevention.feedback

用户通过这个接口对垃圾注册防控结果进行反馈
*/
type AlibabaSecurityJaqSpamregisterpreventionFeedbackRequest struct {
    model.Params

    // 查询接口返回的id
    itemId   string 

    // 聚安全返回的决定信息。当feedBack为0时可以不添
    jaqDecision   int64 

    // 用户自己的决定信息。当feedBack为0时可以不添。
    customerDecision   int64 

    // 用户不认同聚安全返回结果的原因类型。0：同意判定；1：和三方结果不符；2：用户投诉；3:经过人工review判断;9:	其他。
    feedBack   int64 

    // 用户不认同聚安全返回结果的原因描述
    denyReason   string 

}

func NewAlibabaSecurityJaqSpamregisterpreventionFeedbackRequest() *AlibabaSecurityJaqSpamregisterpreventionFeedbackRequest{
    return &AlibabaSecurityJaqSpamregisterpreventionFeedbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqSpamregisterpreventionFeedbackRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.spamregisterprevention.feedback"
}

func (r AlibabaSecurityJaqSpamregisterpreventionFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqSpamregisterpreventionFeedbackRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionFeedbackRequest) GetItemId() string {
    return r.itemId
}

func (r *AlibabaSecurityJaqSpamregisterpreventionFeedbackRequest) SetJaqDecision(jaqDecision int64) error {
    r.jaqDecision = jaqDecision
    r.Set("jaq_decision", jaqDecision)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionFeedbackRequest) GetJaqDecision() int64 {
    return r.jaqDecision
}

func (r *AlibabaSecurityJaqSpamregisterpreventionFeedbackRequest) SetCustomerDecision(customerDecision int64) error {
    r.customerDecision = customerDecision
    r.Set("customer_decision", customerDecision)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionFeedbackRequest) GetCustomerDecision() int64 {
    return r.customerDecision
}

func (r *AlibabaSecurityJaqSpamregisterpreventionFeedbackRequest) SetFeedBack(feedBack int64) error {
    r.feedBack = feedBack
    r.Set("feed_back", feedBack)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionFeedbackRequest) GetFeedBack() int64 {
    return r.feedBack
}

func (r *AlibabaSecurityJaqSpamregisterpreventionFeedbackRequest) SetDenyReason(denyReason string) error {
    r.denyReason = denyReason
    r.Set("deny_reason", denyReason)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionFeedbackRequest) GetDenyReason() string {
    return r.denyReason
}

