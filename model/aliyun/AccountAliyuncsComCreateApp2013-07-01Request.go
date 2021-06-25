package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
给指定用户创建appkey APIRequest
account.aliyuncs.com.CreateApp.2013-07-01

为某个用户创建appkey
*/
type AccountAliyuncsComCreateApp2013-07-01Request struct {
    model.Params

}

func NewAccountAliyuncsComCreateApp2013-07-01Request() *AccountAliyuncsComCreateApp2013-07-01Request{
    return &AccountAliyuncsComCreateApp2013-07-01Request{
        Params: model.NewParams(),
    }
}

func (r AccountAliyuncsComCreateApp2013-07-01Request) GetApiMethodName() string {
    return "account.aliyuncs.com.CreateApp.2013-07-01"
}

func (r AccountAliyuncsComCreateApp2013-07-01Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


