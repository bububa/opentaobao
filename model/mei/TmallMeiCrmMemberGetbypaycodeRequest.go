package mei

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付码获取会员信息 APIRequest
tmall.mei.crm.member.getbypaycode

通过支付码获取会员信息
*/
type TmallMeiCrmMemberGetbypaycodeRequest struct {
    model.Params

    // 会员码
    payCode   string 

}

func NewTmallMeiCrmMemberGetbypaycodeRequest() *TmallMeiCrmMemberGetbypaycodeRequest{
    return &TmallMeiCrmMemberGetbypaycodeRequest{
        Params: model.NewParams(),
    }
}

func (r TmallMeiCrmMemberGetbypaycodeRequest) GetApiMethodName() string {
    return "tmall.mei.crm.member.getbypaycode"
}

func (r TmallMeiCrmMemberGetbypaycodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallMeiCrmMemberGetbypaycodeRequest) SetPayCode(payCode string) error {
    r.payCode = payCode
    r.Set("pay_code", payCode)
    return nil
}

func (r TmallMeiCrmMemberGetbypaycodeRequest) GetPayCode() string {
    return r.payCode
}

