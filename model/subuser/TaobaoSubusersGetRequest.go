package subuser

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定账户的子账号简易信息列表 APIRequest
taobao.subusers.get

获取主账号下的子账号简易账号信息集合。（只能通过主账号登陆并且查询该属于主账号的子账号信息）
*/
type TaobaoSubusersGetRequest struct {
    model.Params

    // 主账号用户名
    userNick   string 

}

func NewTaobaoSubusersGetRequest() *TaobaoSubusersGetRequest{
    return &TaobaoSubusersGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSubusersGetRequest) GetApiMethodName() string {
    return "taobao.subusers.get"
}

func (r TaobaoSubusersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSubusersGetRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

func (r TaobaoSubusersGetRequest) GetUserNick() string {
    return r.userNick
}

