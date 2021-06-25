package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判断买家是否有某些标 APIRequest
taobao.qianniu.buyer.tag.get

判断某个买家是否有某些标
*/
type TaobaoQianniuBuyerTagGetRequest struct {
    model.Params

    // 买家nick
    buyerNick   string 

    // 支持的表，多个tag用英文逗号切割
    tagList   string 

}

func NewTaobaoQianniuBuyerTagGetRequest() *TaobaoQianniuBuyerTagGetRequest{
    return &TaobaoQianniuBuyerTagGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQianniuBuyerTagGetRequest) GetApiMethodName() string {
    return "taobao.qianniu.buyer.tag.get"
}

func (r TaobaoQianniuBuyerTagGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQianniuBuyerTagGetRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r TaobaoQianniuBuyerTagGetRequest) GetBuyerNick() string {
    return r.buyerNick
}

func (r *TaobaoQianniuBuyerTagGetRequest) SetTagList(tagList string) error {
    r.tagList = tagList
    r.Set("tag_list", tagList)
    return nil
}

func (r TaobaoQianniuBuyerTagGetRequest) GetTagList() string {
    return r.tagList
}

