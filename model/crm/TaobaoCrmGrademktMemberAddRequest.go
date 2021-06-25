package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-会员吸纳 APIRequest
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

func NewTaobaoCrmGrademktMemberAddRequest() *TaobaoCrmGrademktMemberAddRequest{
    return &TaobaoCrmGrademktMemberAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmGrademktMemberAddRequest) GetApiMethodName() string {
    return "taobao.crm.grademkt.member.add"
}

func (r TaobaoCrmGrademktMemberAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmGrademktMemberAddRequest) SetParameter(parameter string) error {
    r.parameter = parameter
    r.Set("parameter", parameter)
    return nil
}

func (r TaobaoCrmGrademktMemberAddRequest) GetParameter() string {
    return r.parameter
}

func (r *TaobaoCrmGrademktMemberAddRequest) SetFeather(feather string) error {
    r.feather = feather
    r.Set("feather", feather)
    return nil
}

func (r TaobaoCrmGrademktMemberAddRequest) GetFeather() string {
    return r.feather
}

func (r *TaobaoCrmGrademktMemberAddRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r TaobaoCrmGrademktMemberAddRequest) GetBuyerNick() string {
    return r.buyerNick
}

