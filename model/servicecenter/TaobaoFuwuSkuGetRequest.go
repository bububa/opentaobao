package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取内购服务及SKU详情 APIRequest
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

func NewTaobaoFuwuSkuGetRequest() *TaobaoFuwuSkuGetRequest{
    return &TaobaoFuwuSkuGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFuwuSkuGetRequest) GetApiMethodName() string {
    return "taobao.fuwu.sku.get"
}

func (r TaobaoFuwuSkuGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFuwuSkuGetRequest) SetArticleCode(articleCode string) error {
    r.articleCode = articleCode
    r.Set("article_code", articleCode)
    return nil
}

func (r TaobaoFuwuSkuGetRequest) GetArticleCode() string {
    return r.articleCode
}

func (r *TaobaoFuwuSkuGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoFuwuSkuGetRequest) GetNick() string {
    return r.nick
}

