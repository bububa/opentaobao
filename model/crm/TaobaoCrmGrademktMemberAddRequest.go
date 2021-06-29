package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-会员吸纳 API请求
taobao.crm.grademkt.member.add

商家通过该接口吸纳线上店铺会员。
*/
type TaobaoCrmGrademktMemberAddRequest struct {
    model.Params
    // 会员属性-json format生成方法见http://open.taobao.com/doc/detail.htm?id=101281
    parameter   string
    // 系统属性，json格式
    feather   string
    // 会员nick
    buyerNick   string
}

// 初始化TaobaoCrmGrademktMemberAddRequest对象
func NewTaobaoCrmGrademktMemberAddRequest() *TaobaoCrmGrademktMemberAddRequest{
    return &TaobaoCrmGrademktMemberAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmGrademktMemberAddRequest) GetApiMethodName() string {
    return "taobao.crm.grademkt.member.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmGrademktMemberAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Parameter Setter
// 会员属性-json format生成方法见http://open.taobao.com/doc/detail.htm?id=101281
func (r *TaobaoCrmGrademktMemberAddRequest) SetParameter(parameter string) error {
    r.parameter = parameter
    r.Set("parameter", parameter)
    return nil
}

// Parameter Getter
func (r TaobaoCrmGrademktMemberAddRequest) GetParameter() string {
    return r.parameter
}
// Feather Setter
// 系统属性，json格式
func (r *TaobaoCrmGrademktMemberAddRequest) SetFeather(feather string) error {
    r.feather = feather
    r.Set("feather", feather)
    return nil
}

// Feather Getter
func (r TaobaoCrmGrademktMemberAddRequest) GetFeather() string {
    return r.feather
}
// BuyerNick Setter
// 会员nick
func (r *TaobaoCrmGrademktMemberAddRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoCrmGrademktMemberAddRequest) GetBuyerNick() string {
    return r.buyerNick
}
