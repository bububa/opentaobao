package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-会员关系查询 API请求
taobao.crm.grademkt.member.query

商家通过该接口查询线上店铺会员。
*/
type TaobaoCrmGrademktMemberQueryAPIRequest struct {
    model.Params
    // 会员属性-json format生成方法见http://open.taobao.com/doc/detail.htm?id=101281
    _parameter   string
    // 系统属性，json格式
    _feather   string
    // 会员nick
    _buyerNick   string
}

// 初始化TaobaoCrmGrademktMemberQueryAPIRequest对象
func NewTaobaoCrmGrademktMemberQueryRequest() *TaobaoCrmGrademktMemberQueryAPIRequest{
    return &TaobaoCrmGrademktMemberQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmGrademktMemberQueryAPIRequest) GetApiMethodName() string {
    return "taobao.crm.grademkt.member.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmGrademktMemberQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Parameter Setter
// 会员属性-json format生成方法见http://open.taobao.com/doc/detail.htm?id=101281
func (r *TaobaoCrmGrademktMemberQueryAPIRequest) SetParameter(_parameter string) error {
    r._parameter = _parameter
    r.Set("parameter", _parameter)
    return nil
}

// Parameter Getter
func (r TaobaoCrmGrademktMemberQueryAPIRequest) GetParameter() string {
    return r._parameter
}
// Feather Setter
// 系统属性，json格式
func (r *TaobaoCrmGrademktMemberQueryAPIRequest) SetFeather(_feather string) error {
    r._feather = _feather
    r.Set("feather", _feather)
    return nil
}

// Feather Getter
func (r TaobaoCrmGrademktMemberQueryAPIRequest) GetFeather() string {
    return r._feather
}
// BuyerNick Setter
// 会员nick
func (r *TaobaoCrmGrademktMemberQueryAPIRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoCrmGrademktMemberQueryAPIRequest) GetBuyerNick() string {
    return r._buyerNick
}
