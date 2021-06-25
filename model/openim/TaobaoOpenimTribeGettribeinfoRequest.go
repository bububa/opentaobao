package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取群信息 APIRequest
taobao.openim.tribe.gettribeinfo

获取群信息
*/
type TaobaoOpenimTribeGettribeinfoRequest struct {
    model.Params

    // 用户信息
    user   *OpenImUser 

    // 群ID
    tribeId   int64 

}

func NewTaobaoOpenimTribeGettribeinfoRequest() *TaobaoOpenimTribeGettribeinfoRequest{
    return &TaobaoOpenimTribeGettribeinfoRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimTribeGettribeinfoRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.gettribeinfo"
}

func (r TaobaoOpenimTribeGettribeinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimTribeGettribeinfoRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

func (r TaobaoOpenimTribeGettribeinfoRequest) GetUser() *OpenImUser {
    return r.user
}

func (r *TaobaoOpenimTribeGettribeinfoRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

func (r TaobaoOpenimTribeGettribeinfoRequest) GetTribeId() int64 {
    return r.tribeId
}

