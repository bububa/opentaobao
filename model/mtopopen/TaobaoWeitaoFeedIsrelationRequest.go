package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
是否关注 APIRequest
taobao.weitao.feed.isrelation

判断用户是否关注对应的公共账号
*/
type TaobaoWeitaoFeedIsrelationRequest struct {
    model.Params

    // 要查询的粉丝的淘宝昵称
    fansNick   string 

    // 要查询的公共账号的淘宝昵称
    sellerNick   string 

}

func NewTaobaoWeitaoFeedIsrelationRequest() *TaobaoWeitaoFeedIsrelationRequest{
    return &TaobaoWeitaoFeedIsrelationRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWeitaoFeedIsrelationRequest) GetApiMethodName() string {
    return "taobao.weitao.feed.isrelation"
}

func (r TaobaoWeitaoFeedIsrelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWeitaoFeedIsrelationRequest) SetFansNick(fansNick string) error {
    r.fansNick = fansNick
    r.Set("fans_nick", fansNick)
    return nil
}

func (r TaobaoWeitaoFeedIsrelationRequest) GetFansNick() string {
    return r.fansNick
}

func (r *TaobaoWeitaoFeedIsrelationRequest) SetSellerNick(sellerNick string) error {
    r.sellerNick = sellerNick
    r.Set("seller_nick", sellerNick)
    return nil
}

func (r TaobaoWeitaoFeedIsrelationRequest) GetSellerNick() string {
    return r.sellerNick
}

