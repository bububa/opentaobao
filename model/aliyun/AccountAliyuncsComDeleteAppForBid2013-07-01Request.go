package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
运营商删除用户的appkey APIRequest
account.aliyuncs.com.DeleteAppForBid.2013-07-01

删除用户的appkey，会校验调用的用户是否为当前运营商所创建的。
*/
type AccountAliyuncsComDeleteAppForBid2013-07-01Request struct {
    model.Params

    // 要删除的appkey的所有者用户的pk
    ownerId   string 

    // 要删除的appkey
    ownerAppkey   string 

}

func NewAccountAliyuncsComDeleteAppForBid2013-07-01Request() *AccountAliyuncsComDeleteAppForBid2013-07-01Request{
    return &AccountAliyuncsComDeleteAppForBid2013-07-01Request{
        Params: model.NewParams(),
    }
}

func (r AccountAliyuncsComDeleteAppForBid2013-07-01Request) GetApiMethodName() string {
    return "account.aliyuncs.com.DeleteAppForBid.2013-07-01"
}

func (r AccountAliyuncsComDeleteAppForBid2013-07-01Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AccountAliyuncsComDeleteAppForBid2013-07-01Request) SetOwnerId(ownerId string) error {
    r.ownerId = ownerId
    r.Set("OwnerId", ownerId)
    return nil
}

func (r AccountAliyuncsComDeleteAppForBid2013-07-01Request) GetOwnerId() string {
    return r.ownerId
}

func (r *AccountAliyuncsComDeleteAppForBid2013-07-01Request) SetOwnerAppkey(ownerAppkey string) error {
    r.ownerAppkey = ownerAppkey
    r.Set("OwnerAppkey", ownerAppkey)
    return nil
}

func (r AccountAliyuncsComDeleteAppForBid2013-07-01Request) GetOwnerAppkey() string {
    return r.ownerAppkey
}

