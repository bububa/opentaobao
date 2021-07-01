package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取登陆权限签名 API请求
taobao.simba.login.authsign.get

获取登陆权限签名
*/
type TaobaoSimbaLoginAuthsignGetAPIRequest struct {
    model.Params
    // 主人昵称
    _nick   string
}

// 初始化TaobaoSimbaLoginAuthsignGetAPIRequest对象
func NewTaobaoSimbaLoginAuthsignGetRequest() *TaobaoSimbaLoginAuthsignGetAPIRequest{
    return &TaobaoSimbaLoginAuthsignGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaLoginAuthsignGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.login.authsign.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaLoginAuthsignGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaLoginAuthsignGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaLoginAuthsignGetAPIRequest) GetNick() string {
    return r._nick
}
