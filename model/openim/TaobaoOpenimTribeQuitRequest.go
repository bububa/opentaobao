package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群成员退出 APIRequest
taobao.openim.tribe.quit

OPENIM群成员退出
*/
type TaobaoOpenimTribeQuitRequest struct {
    model.Params

    // 用户信息
    user   *OpenImUser 

    // 群id
    tribeId   int64 

}

func NewTaobaoOpenimTribeQuitRequest() *TaobaoOpenimTribeQuitRequest{
    return &TaobaoOpenimTribeQuitRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimTribeQuitRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.quit"
}

func (r TaobaoOpenimTribeQuitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimTribeQuitRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

func (r TaobaoOpenimTribeQuitRequest) GetUser() *OpenImUser {
    return r.user
}

func (r *TaobaoOpenimTribeQuitRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

func (r TaobaoOpenimTribeQuitRequest) GetTribeId() int64 {
    return r.tribeId
}

