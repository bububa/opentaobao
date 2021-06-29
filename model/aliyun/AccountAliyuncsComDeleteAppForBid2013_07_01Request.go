package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
运营商删除用户的appkey API请求
account.aliyuncs.com.DeleteAppForBid.2013-07-01

删除用户的appkey，会校验调用的用户是否为当前运营商所创建的。
*/
type AccountAliyuncsComDeleteAppForBid2013_07_01Request struct {
    model.Params
    // 要删除的appkey的所有者用户的pk
    _ownerId   string
    // 要删除的appkey
    _ownerAppkey   string
}

// 初始化AccountAliyuncsComDeleteAppForBid2013_07_01Request对象
func NewAccountAliyuncsComDeleteAppForBid2013_07_01Request() *AccountAliyuncsComDeleteAppForBid2013_07_01Request{
    return &AccountAliyuncsComDeleteAppForBid2013_07_01Request{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComDeleteAppForBid2013_07_01Request) GetApiMethodName() string {
    return "account.aliyuncs.com.DeleteAppForBid.2013-07-01"
}

// IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComDeleteAppForBid2013_07_01Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OwnerId Setter
// 要删除的appkey的所有者用户的pk
func (r *AccountAliyuncsComDeleteAppForBid2013_07_01Request) SetOwnerId(_ownerId string) error {
    r._ownerId = _ownerId
    r.Set("OwnerId", _ownerId)
    return nil
}

// OwnerId Getter
func (r AccountAliyuncsComDeleteAppForBid2013_07_01Request) GetOwnerId() string {
    return r._ownerId
}
// OwnerAppkey Setter
// 要删除的appkey
func (r *AccountAliyuncsComDeleteAppForBid2013_07_01Request) SetOwnerAppkey(_ownerAppkey string) error {
    r._ownerAppkey = _ownerAppkey
    r.Set("OwnerAppkey", _ownerAppkey)
    return nil
}

// OwnerAppkey Getter
func (r AccountAliyuncsComDeleteAppForBid2013_07_01Request) GetOwnerAppkey() string {
    return r._ownerAppkey
}