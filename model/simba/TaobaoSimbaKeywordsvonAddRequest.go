package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建一批关键词 API请求
taobao.simba.keywordsvon.add

创建一批关键词
*/
type TaobaoSimbaKeywordsvonAddRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 推广组id
    _adgroupId   int64
    // 关键词、计算机出价、无线出价和匹配方式json字符串，word:词，不能有一些特殊字符。maxPrice：计算机出价，是整数，以“分”为单位，不能小于5，不能大于日限额, maxMobilePrice：代表无线出价，规则同maxPice 当matchscope只能是1,4（1代表精确匹配，4代表广泛匹配）。
    _keywordPrices   string
}

// 初始化TaobaoSimbaKeywordsvonAddRequest对象
func NewTaobaoSimbaKeywordsvonAddRequest() *TaobaoSimbaKeywordsvonAddRequest{
    return &TaobaoSimbaKeywordsvonAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsvonAddRequest) GetApiMethodName() string {
    return "taobao.simba.keywordsvon.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsvonAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordsvonAddRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaKeywordsvonAddRequest) GetNick() string {
    return r._nick
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaKeywordsvonAddRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaKeywordsvonAddRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// KeywordPrices Setter
// 关键词、计算机出价、无线出价和匹配方式json字符串，word:词，不能有一些特殊字符。maxPrice：计算机出价，是整数，以“分”为单位，不能小于5，不能大于日限额, maxMobilePrice：代表无线出价，规则同maxPice 当matchscope只能是1,4（1代表精确匹配，4代表广泛匹配）。
func (r *TaobaoSimbaKeywordsvonAddRequest) SetKeywordPrices(_keywordPrices string) error {
    r._keywordPrices = _keywordPrices
    r.Set("keyword_prices", _keywordPrices)
    return nil
}

// KeywordPrices Getter
func (r TaobaoSimbaKeywordsvonAddRequest) GetKeywordPrices() string {
    return r._keywordPrices
}
