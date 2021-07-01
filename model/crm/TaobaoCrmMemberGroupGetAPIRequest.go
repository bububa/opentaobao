package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取买家身上的标签 API请求
taobao.crm.member.group.get

获取买家身上的标签，不返回标签的总人数
*/
type TaobaoCrmMemberGroupGetAPIRequest struct {
    model.Params
    // 会员Nick
    _buyerNick   string
}

// 初始化TaobaoCrmMemberGroupGetAPIRequest对象
func NewTaobaoCrmMemberGroupGetRequest() *TaobaoCrmMemberGroupGetAPIRequest{
    return &TaobaoCrmMemberGroupGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmMemberGroupGetAPIRequest) GetApiMethodName() string {
    return "taobao.crm.member.group.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmMemberGroupGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BuyerNick Setter
// 会员Nick
func (r *TaobaoCrmMemberGroupGetAPIRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoCrmMemberGroupGetAPIRequest) GetBuyerNick() string {
    return r._buyerNick
}
