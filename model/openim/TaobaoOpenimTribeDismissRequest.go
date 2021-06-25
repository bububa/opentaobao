package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群解散 APIRequest
taobao.openim.tribe.dismiss

OPENIM群解散
*/
type TaobaoOpenimTribeDismissRequest struct {
    model.Params

    // 用户信息
    user   *OpenImUser 

    // 群id
    tribeId   int64 

}

func NewTaobaoOpenimTribeDismissRequest() *TaobaoOpenimTribeDismissRequest{
    return &TaobaoOpenimTribeDismissRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimTribeDismissRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.dismiss"
}

func (r TaobaoOpenimTribeDismissRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimTribeDismissRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

func (r TaobaoOpenimTribeDismissRequest) GetUser() *OpenImUser {
    return r.user
}

func (r *TaobaoOpenimTribeDismissRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

func (r TaobaoOpenimTribeDismissRequest) GetTribeId() int64 {
    return r.tribeId
}

