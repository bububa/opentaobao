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
type TmallMeiCrmMemberGetbypaycodeAPIRequest struct {
    model.Params
    // 会员码
    _payCode   string
}

// 初始化TmallMeiCrmMemberGetbypaycodeAPIRequest对象
func NewTmallMeiCrmMemberGetbypaycodeRequest() *TmallMeiCrmMemberGetbypaycodeAPIRequest{
    return &TmallMeiCrmMemberGetbypaycodeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMeiCrmMemberGetbypaycodeAPIRequest) GetApiMethodName() string {
    return "tmall.mei.crm.member.getbypaycode"
}

// IRequest interface 方法, 获取API参数
func (r TmallMeiCrmMemberGetbypaycodeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PayCode Setter
// 会员码
func (r *TmallMeiCrmMemberGetbypaycodeAPIRequest) SetPayCode(_payCode string) error {
    r._payCode = _payCode
    r.Set("pay_code", _payCode)
    return nil
}

// PayCode Getter
func (r TmallMeiCrmMemberGetbypaycodeAPIRequest) GetPayCode() string {
    return r._payCode
}
