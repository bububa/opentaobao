package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取搜索人群TOP用户可添加人群信息 API请求
taobao.simba.searchtagtemplate.get

获取搜索人群用户可添加人群信息
*/
type TaobaoSimbaSearchtagtemplateGetRequest struct {
    model.Params
    // 被操作者的淘宝昵称
    _nick   string
    // 子帐号nick
    _subNick   string
}

// 初始化TaobaoSimbaSearchtagtemplateGetRequest对象
func NewTaobaoSimbaSearchtagtemplateGetRequest() *TaobaoSimbaSearchtagtemplateGetRequest{
    return &TaobaoSimbaSearchtagtemplateGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSearchtagtemplateGetRequest) GetApiMethodName() string {
    return "taobao.simba.searchtagtemplate.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSearchtagtemplateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaoSimbaSearchtagtemplateGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaSearchtagtemplateGetRequest) GetNick() string {
    return r._nick
}
// SubNick Setter
// 子帐号nick
func (r *TaobaoSimbaSearchtagtemplateGetRequest) SetSubNick(_subNick string) error {
    r._subNick = _subNick
    r.Set("sub_nick", _subNick)
    return nil
}

// SubNick Getter
func (r TaobaoSimbaSearchtagtemplateGetRequest) GetSubNick() string {
    return r._subNick
}
