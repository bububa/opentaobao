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
type TaobaoCrmMemberGroupGetRequest struct {
    model.Params
    // 会员Nick
    buyerNick   string
}

// 初始化TaobaoCrmMemberGroupGetRequest对象
func NewTaobaoCrmMemberGroupGetRequest() *TaobaoCrmMemberGroupGetRequest{
    return &TaobaoCrmMemberGroupGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmMemberGroupGetRequest) GetApiMethodName() string {
    return "taobao.crm.member.group.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmMemberGroupGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BuyerNick Setter
// 会员Nick
func (r *TaobaoCrmMemberGroupGetRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoCrmMemberGroupGetRequest) GetBuyerNick() string {
    return r.buyerNick
}
