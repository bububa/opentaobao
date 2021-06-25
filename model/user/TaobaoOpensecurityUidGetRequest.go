package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝open security uid 获取接口 APIRequest
taobao.opensecurity.uid.get

根据明文 taobao user id 换取 app的 open_uid
*/
type TaobaoOpensecurityUidGetRequest struct {
    model.Params

    // 淘宝用户ID
    tbUserId   int64 

}

func NewTaobaoOpensecurityUidGetRequest() *TaobaoOpensecurityUidGetRequest{
    return &TaobaoOpensecurityUidGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpensecurityUidGetRequest) GetApiMethodName() string {
    return "taobao.opensecurity.uid.get"
}

func (r TaobaoOpensecurityUidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpensecurityUidGetRequest) SetTbUserId(tbUserId int64) error {
    r.tbUserId = tbUserId
    r.Set("tb_user_id", tbUserId)
    return nil
}

func (r TaobaoOpensecurityUidGetRequest) GetTbUserId() int64 {
    return r.tbUserId
}

