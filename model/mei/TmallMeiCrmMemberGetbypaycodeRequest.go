package mei

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付码获取会员信息 API请求
tmall.mei.crm.member.getbypaycode

通过支付码获取会员信息
*/
type TmallMeiCrmMemberGetbypaycodeRequest struct {
    model.Params
    // 会员码
    payCode   string
}

// 初始化TmallMeiCrmMemberGetbypaycodeRequest对象
func NewTmallMeiCrmMemberGetbypaycodeRequest() *TmallMeiCrmMemberGetbypaycodeRequest{
    return &TmallMeiCrmMemberGetbypaycodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMeiCrmMemberGetbypaycodeRequest) GetApiMethodName() string {
    return "tmall.mei.crm.member.getbypaycode"
}

// IRequest interface 方法, 获取API参数
func (r TmallMeiCrmMemberGetbypaycodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PayCode Setter
// 会员码
func (r *TmallMeiCrmMemberGetbypaycodeRequest) SetPayCode(payCode string) error {
    r.payCode = payCode
    r.Set("pay_code", payCode)
    return nil
}

// PayCode Getter
func (r TmallMeiCrmMemberGetbypaycodeRequest) GetPayCode() string {
    return r.payCode
}
