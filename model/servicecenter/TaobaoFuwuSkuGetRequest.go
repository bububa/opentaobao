package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取内购服务及SKU详情 API请求
taobao.fuwu.sku.get

通过服务code和用户nick，获取该服务对应的收费项目的sku信息，包括价格、可购买周期、用户能否购买等信息
*/
type TaobaoFuwuSkuGetRequest struct {
    model.Params
    // 服务code
    articleCode   string
    // 用户的淘宝nick
    nick   string
}

// 初始化TaobaoFuwuSkuGetRequest对象
func NewTaobaoFuwuSkuGetRequest() *TaobaoFuwuSkuGetRequest{
    return &TaobaoFuwuSkuGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFuwuSkuGetRequest) GetApiMethodName() string {
    return "taobao.fuwu.sku.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFuwuSkuGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ArticleCode Setter
// 服务code
func (r *TaobaoFuwuSkuGetRequest) SetArticleCode(articleCode string) error {
    r.articleCode = articleCode
    r.Set("article_code", articleCode)
    return nil
}

// ArticleCode Getter
func (r TaobaoFuwuSkuGetRequest) GetArticleCode() string {
    return r.articleCode
}
// Nick Setter
// 用户的淘宝nick
func (r *TaobaoFuwuSkuGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoFuwuSkuGetRequest) GetNick() string {
    return r.nick
}
