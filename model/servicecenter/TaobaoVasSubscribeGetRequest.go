package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订购关系查询 APIRequest
taobao.vas.subscribe.get

用于ISV根据登录进来的淘宝会员名查询该为该会员开通哪些收费项目，ISV只能查询自己名下的应用及收费项目的订购情况
*/
type TaobaoVasSubscribeGetRequest struct {
    model.Params

    // 商品编码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的商品代码
    articleCode   string 

    // 淘宝会员名
    nick   string 

}

func NewTaobaoVasSubscribeGetRequest() *TaobaoVasSubscribeGetRequest{
    return &TaobaoVasSubscribeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVasSubscribeGetRequest) GetApiMethodName() string {
    return "taobao.vas.subscribe.get"
}

func (r TaobaoVasSubscribeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVasSubscribeGetRequest) SetArticleCode(articleCode string) error {
    r.articleCode = articleCode
    r.Set("article_code", articleCode)
    return nil
}

func (r TaobaoVasSubscribeGetRequest) GetArticleCode() string {
    return r.articleCode
}

func (r *TaobaoVasSubscribeGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoVasSubscribeGetRequest) GetNick() string {
    return r.nick
}

