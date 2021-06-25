package subuser

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定账户的所有职务信息列表 APIRequest
taobao.subuser.dutys.get

通过主账号Nick获取该账户下的所有职务信息，职务信息中包括职务ID、职务名称以及职务等级（通过主账号登陆只能获取属于该主账号下的职务信息）
*/
type TaobaoSubuserDutysGetRequest struct {
    model.Params

    // 主账号用户名
    userNick   string 

}

func NewTaobaoSubuserDutysGetRequest() *TaobaoSubuserDutysGetRequest{
    return &TaobaoSubuserDutysGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSubuserDutysGetRequest) GetApiMethodName() string {
    return "taobao.subuser.dutys.get"
}

func (r TaobaoSubuserDutysGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSubuserDutysGetRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

func (r TaobaoSubuserDutysGetRequest) GetUserNick() string {
    return r.userNick
}

