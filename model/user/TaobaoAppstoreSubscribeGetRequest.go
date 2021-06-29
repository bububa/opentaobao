package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询appstore应用订购关系 API请求
taobao.appstore.subscribe.get

查询appstore应用订购关系(对于新上架的多版本应用，建议使用taobao.vas.subscribe.get)
*/
type TaobaoAppstoreSubscribeGetRequest struct {
    model.Params
    // 用户昵称
    nick   string
}

// 初始化TaobaoAppstoreSubscribeGetRequest对象
func NewTaobaoAppstoreSubscribeGetRequest() *TaobaoAppstoreSubscribeGetRequest{
    return &TaobaoAppstoreSubscribeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAppstoreSubscribeGetRequest) GetApiMethodName() string {
    return "taobao.appstore.subscribe.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAppstoreSubscribeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 用户昵称
func (r *TaobaoAppstoreSubscribeGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoAppstoreSubscribeGetRequest) GetNick() string {
    return r.nick
}
