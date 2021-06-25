package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
为bid用户创建账号 APIRequest
account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01

给指定的bid创建账号，同时账号打上owner bid的标记
*/
type AccountAliyuncsComCreateAliyunAccountForBid2013-07-01Request struct {
    model.Params

}

func NewAccountAliyuncsComCreateAliyunAccountForBid2013-07-01Request() *AccountAliyuncsComCreateAliyunAccountForBid2013-07-01Request{
    return &AccountAliyuncsComCreateAliyunAccountForBid2013-07-01Request{
        Params: model.NewParams(),
    }
}

func (r AccountAliyuncsComCreateAliyunAccountForBid2013-07-01Request) GetApiMethodName() string {
    return "account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01"
}

func (r AccountAliyuncsComCreateAliyunAccountForBid2013-07-01Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


