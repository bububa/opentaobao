package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-会员关系查询 APIRequest
taobao.crm.grademkt.member.query

商家通过该接口查询线上店铺会员。
*/
type TaobaoCrmGrademktMemberQueryRequest struct {
    model.Params

    // 会员属性-json format生成方法见http://open.taobao.com/doc/detail.htm?id=101281
    parameter   string 

    // 系统属性，json格式
    feather   string 

    // 会员nick
    buyerNick   string 

}

func NewTaobaoCrmGrademktMemberQueryRequest() *TaobaoCrmGrademktMemberQueryRequest{
    return &TaobaoCrmGrademktMemberQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmGrademktMemberQueryRequest) GetApiMethodName() string {
    return "taobao.crm.grademkt.member.query"
}

func (r TaobaoCrmGrademktMemberQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmGrademktMemberQueryRequest) SetParameter(parameter string) error {
    r.parameter = parameter
    r.Set("parameter", parameter)
    return nil
}

func (r TaobaoCrmGrademktMemberQueryRequest) GetParameter() string {
    return r.parameter
}

func (r *TaobaoCrmGrademktMemberQueryRequest) SetFeather(feather string) error {
    r.feather = feather
    r.Set("feather", feather)
    return nil
}

func (r TaobaoCrmGrademktMemberQueryRequest) GetFeather() string {
    return r.feather
}

func (r *TaobaoCrmGrademktMemberQueryRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r TaobaoCrmGrademktMemberQueryRequest) GetBuyerNick() string {
    return r.buyerNick
}

