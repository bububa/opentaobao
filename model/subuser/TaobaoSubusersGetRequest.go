package subuser

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定账户的子账号简易信息列表 API请求
taobao.subusers.get

获取主账号下的子账号简易账号信息集合。（只能通过主账号登陆并且查询该属于主账号的子账号信息）
*/
type TaobaoSubusersGetRequest struct {
    model.Params
    // 主账号用户名
    _userNick   string
}

// 初始化TaobaoSubusersGetRequest对象
func NewTaobaoSubusersGetRequest() *TaobaoSubusersGetRequest{
    return &TaobaoSubusersGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSubusersGetRequest) GetApiMethodName() string {
    return "taobao.subusers.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSubusersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserNick Setter
// 主账号用户名
func (r *TaobaoSubusersGetRequest) SetUserNick(_userNick string) error {
    r._userNick = _userNick
    r.Set("user_nick", _userNick)
    return nil
}

// UserNick Getter
func (r TaobaoSubusersGetRequest) GetUserNick() string {
    return r._userNick
}
