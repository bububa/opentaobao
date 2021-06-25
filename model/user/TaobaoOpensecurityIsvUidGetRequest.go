package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取open security uid for isv APIRequest
taobao.opensecurity.isv.uid.get

根据 open_uid 获取 open_uid_isv 用于同一个 isv的多个app间数据关联
*/
type TaobaoOpensecurityIsvUidGetRequest struct {
    model.Params

    // 基于Appkey生成的open security淘宝用户Id
    openUid   string 

}

func NewTaobaoOpensecurityIsvUidGetRequest() *TaobaoOpensecurityIsvUidGetRequest{
    return &TaobaoOpensecurityIsvUidGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpensecurityIsvUidGetRequest) GetApiMethodName() string {
    return "taobao.opensecurity.isv.uid.get"
}

func (r TaobaoOpensecurityIsvUidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpensecurityIsvUidGetRequest) SetOpenUid(openUid string) error {
    r.openUid = openUid
    r.Set("open_uid", openUid)
    return nil
}

func (r TaobaoOpensecurityIsvUidGetRequest) GetOpenUid() string {
    return r.openUid
}

