package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一组推广计划 API请求
taobao.simba.campaigns.get

取得一个客户的推广计划；
*/
type TaobaoSimbaCampaignsGetRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 计划类型0位标准计划，16位销量明星计划
    _type   int64
}

// 初始化TaobaoSimbaCampaignsGetRequest对象
func NewTaobaoSimbaCampaignsGetRequest() *TaobaoSimbaCampaignsGetRequest{
    return &TaobaoSimbaCampaignsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignsGetRequest) GetApiMethodName() string {
    return "taobao.simba.campaigns.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignsGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCampaignsGetRequest) GetNick() string {
    return r._nick
}
// Type Setter
// 计划类型0位标准计划，16位销量明星计划
func (r *TaobaoSimbaCampaignsGetRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoSimbaCampaignsGetRequest) GetType() int64 {
    return r._type
}
