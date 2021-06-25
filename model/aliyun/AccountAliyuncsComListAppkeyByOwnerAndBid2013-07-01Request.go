package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据渠道id和状态列出appkey APIRequest
account.aliyuncs.com.ListAppkeyByOwnerAndBid.2013-07-01

根据渠道id和状态列出appkey
*/
type AccountAliyuncsComListAppkeyByOwnerAndBid2013-07-01Request struct {
    model.Params

}

func NewAccountAliyuncsComListAppkeyByOwnerAndBid2013-07-01Request() *AccountAliyuncsComListAppkeyByOwnerAndBid2013-07-01Request{
    return &AccountAliyuncsComListAppkeyByOwnerAndBid2013-07-01Request{
        Params: model.NewParams(),
    }
}

func (r AccountAliyuncsComListAppkeyByOwnerAndBid2013-07-01Request) GetApiMethodName() string {
    return "account.aliyuncs.com.ListAppkeyByOwnerAndBid.2013-07-01"
}

func (r AccountAliyuncsComListAppkeyByOwnerAndBid2013-07-01Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


