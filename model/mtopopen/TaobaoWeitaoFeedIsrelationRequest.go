package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
是否关注 API请求
taobao.weitao.feed.isrelation

判断用户是否关注对应的公共账号
*/
type TaobaoWeitaoFeedIsrelationRequest struct {
    model.Params
    // 要查询的粉丝的淘宝昵称
    _fansNick   string
    // 要查询的公共账号的淘宝昵称
    _sellerNick   string
}

// 初始化TaobaoWeitaoFeedIsrelationRequest对象
func NewTaobaoWeitaoFeedIsrelationRequest() *TaobaoWeitaoFeedIsrelationRequest{
    return &TaobaoWeitaoFeedIsrelationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWeitaoFeedIsrelationRequest) GetApiMethodName() string {
    return "taobao.weitao.feed.isrelation"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWeitaoFeedIsrelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FansNick Setter
// 要查询的粉丝的淘宝昵称
func (r *TaobaoWeitaoFeedIsrelationRequest) SetFansNick(_fansNick string) error {
    r._fansNick = _fansNick
    r.Set("fans_nick", _fansNick)
    return nil
}

// FansNick Getter
func (r TaobaoWeitaoFeedIsrelationRequest) GetFansNick() string {
    return r._fansNick
}
// SellerNick Setter
// 要查询的公共账号的淘宝昵称
func (r *TaobaoWeitaoFeedIsrelationRequest) SetSellerNick(_sellerNick string) error {
    r._sellerNick = _sellerNick
    r.Set("seller_nick", _sellerNick)
    return nil
}

// SellerNick Getter
func (r TaobaoWeitaoFeedIsrelationRequest) GetSellerNick() string {
    return r._sellerNick
}
