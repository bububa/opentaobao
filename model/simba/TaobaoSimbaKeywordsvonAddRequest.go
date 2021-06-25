package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建一批关键词 APIRequest
taobao.simba.keywordsvon.add

创建一批关键词
*/
type TaobaoSimbaKeywordsvonAddRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 推广组id
    adgroupId   int64 

    // 关键词、计算机出价、无线出价和匹配方式json字符串，word:词，不能有一些特殊字符。maxPrice：计算机出价，是整数，以“分”为单位，不能小于5，不能大于日限额, maxMobilePrice：代表无线出价，规则同maxPice 当matchscope只能是1,4（1代表精确匹配，4代表广泛匹配）。
    keywordPrices   string 

}

func NewTaobaoSimbaKeywordsvonAddRequest() *TaobaoSimbaKeywordsvonAddRequest{
    return &TaobaoSimbaKeywordsvonAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaKeywordsvonAddRequest) GetApiMethodName() string {
    return "taobao.simba.keywordsvon.add"
}

func (r TaobaoSimbaKeywordsvonAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaKeywordsvonAddRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaKeywordsvonAddRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaKeywordsvonAddRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaKeywordsvonAddRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaKeywordsvonAddRequest) SetKeywordPrices(keywordPrices string) error {
    r.keywordPrices = keywordPrices
    r.Set("keyword_prices", keywordPrices)
    return nil
}

func (r TaobaoSimbaKeywordsvonAddRequest) GetKeywordPrices() string {
    return r.keywordPrices
}

