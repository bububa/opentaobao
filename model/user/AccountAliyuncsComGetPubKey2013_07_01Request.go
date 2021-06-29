package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户公钥 APIRequest
account.aliyuncs.com.GetPubKey.2013-07-01

根据用户的appkey查询用户的pubkey
*/
type AccountAliyuncsComGetPubKey2013_07_01Request struct {
    model.Params

    // appkey
    ownerAppkey   string 

}

func NewAccountAliyuncsComGetPubKey2013_07_01Request() *AccountAliyuncsComGetPubKey2013_07_01Request{
    return &AccountAliyuncsComGetPubKey2013_07_01Request{
        Params: model.NewParams(),
    }
}

func (r AccountAliyuncsComGetPubKey2013_07_01Request) GetApiMethodName() string {
    return "account.aliyuncs.com.GetPubKey.2013-07-01"
}

func (r AccountAliyuncsComGetPubKey2013_07_01Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AccountAliyuncsComGetPubKey2013_07_01Request) SetOwnerAppkey(ownerAppkey string) error {
    r.ownerAppkey = ownerAppkey
    r.Set("OwnerAppkey", ownerAppkey)
    return nil
}

func (r AccountAliyuncsComGetPubKey2013_07_01Request) GetOwnerAppkey() string {
    return r.ownerAppkey
}

