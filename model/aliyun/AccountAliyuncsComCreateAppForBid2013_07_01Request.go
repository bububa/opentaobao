package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
给当前渠道下的用户创建app APIRequest
account.aliyuncs.com.CreateAppForBid.2013-07-01

给自己渠道下的用户创建app
*/
type AccountAliyuncsComCreateAppForBid2013_07_01Request struct {
    model.Params

}

func NewAccountAliyuncsComCreateAppForBid2013_07_01Request() *AccountAliyuncsComCreateAppForBid2013_07_01Request{
    return &AccountAliyuncsComCreateAppForBid2013_07_01Request{
        Params: model.NewParams(),
    }
}

func (r AccountAliyuncsComCreateAppForBid2013_07_01Request) GetApiMethodName() string {
    return "account.aliyuncs.com.CreateAppForBid.2013-07-01"
}

func (r AccountAliyuncsComCreateAppForBid2013_07_01Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


