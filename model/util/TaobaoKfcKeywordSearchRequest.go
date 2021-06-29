package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词过滤匹配 API请求
taobao.kfc.keyword.search

对输入的文本信息进行禁忌关键词匹配，返回匹配的结果
*/
type TaobaoKfcKeywordSearchRequest struct {
    model.Params
    // 发布信息的淘宝会员名，可以不传
    _nick   string
    // 需要过滤的文本信息
    _content   string
    // 应用点，分为一级应用点、二级应用点。其中一级应用点通常是指某一个系统或产品，比如淘宝的商品应用（taobao_auction）；二级应用点，是指一级应用点下的具体的分类，比如商品标题(title)、商品描述(content)。不同的二级应用可以设置不同关键词。<br/><br/>这里的apply参数是由一级应用点与二级应用点合起来的字符（一级应用点+"."+二级应用点），如taobao_auction.title。<br/><br/><br/>通常apply参数是不需要传递的。如有特殊需求（比如特殊的过滤需求，需要自己维护一套自己词库），需传递此参数。
    _apply   string
}

// 初始化TaobaoKfcKeywordSearchRequest对象
func NewTaobaoKfcKeywordSearchRequest() *TaobaoKfcKeywordSearchRequest{
    return &TaobaoKfcKeywordSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKfcKeywordSearchRequest) GetApiMethodName() string {
    return "taobao.kfc.keyword.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKfcKeywordSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 发布信息的淘宝会员名，可以不传
func (r *TaobaoKfcKeywordSearchRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoKfcKeywordSearchRequest) GetNick() string {
    return r._nick
}
// Content Setter
// 需要过滤的文本信息
func (r *TaobaoKfcKeywordSearchRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r TaobaoKfcKeywordSearchRequest) GetContent() string {
    return r._content
}
// Apply Setter
// 应用点，分为一级应用点、二级应用点。其中一级应用点通常是指某一个系统或产品，比如淘宝的商品应用（taobao_auction）；二级应用点，是指一级应用点下的具体的分类，比如商品标题(title)、商品描述(content)。不同的二级应用可以设置不同关键词。<br/><br/>这里的apply参数是由一级应用点与二级应用点合起来的字符（一级应用点+"."+二级应用点），如taobao_auction.title。<br/><br/><br/>通常apply参数是不需要传递的。如有特殊需求（比如特殊的过滤需求，需要自己维护一套自己词库），需传递此参数。
func (r *TaobaoKfcKeywordSearchRequest) SetApply(_apply string) error {
    r._apply = _apply
    r.Set("apply", _apply)
    return nil
}

// Apply Getter
func (r TaobaoKfcKeywordSearchRequest) GetApply() string {
    return r._apply
}
