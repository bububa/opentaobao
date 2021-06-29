package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判断买家是否有某些标 API请求
taobao.qianniu.buyer.tag.get

判断某个买家是否有某些标
*/
type TaobaoQianniuBuyerTagGetRequest struct {
    model.Params
    // 买家nick
    _buyerNick   string
    // 支持的表，多个tag用英文逗号切割
    _tagList   string
}

// 初始化TaobaoQianniuBuyerTagGetRequest对象
func NewTaobaoQianniuBuyerTagGetRequest() *TaobaoQianniuBuyerTagGetRequest{
    return &TaobaoQianniuBuyerTagGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQianniuBuyerTagGetRequest) GetApiMethodName() string {
    return "taobao.qianniu.buyer.tag.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQianniuBuyerTagGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BuyerNick Setter
// 买家nick
func (r *TaobaoQianniuBuyerTagGetRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoQianniuBuyerTagGetRequest) GetBuyerNick() string {
    return r._buyerNick
}
// TagList Setter
// 支持的表，多个tag用英文逗号切割
func (r *TaobaoQianniuBuyerTagGetRequest) SetTagList(_tagList string) error {
    r._tagList = _tagList
    r.Set("tag_list", _tagList)
    return nil
}

// TagList Getter
func (r TaobaoQianniuBuyerTagGetRequest) GetTagList() string {
    return r._tagList
}
