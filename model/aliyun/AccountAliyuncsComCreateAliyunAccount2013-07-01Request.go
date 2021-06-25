package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建阿里云账号 APIRequest
account.aliyuncs.com.CreateAliyunAccount.2013-07-01

根据给定的阿里云账号，密码以及手机号创建阿里云账号
*/
type AccountAliyuncsComCreateAliyunAccount2013-07-01Request struct {
    model.Params

}

func NewAccountAliyuncsComCreateAliyunAccount2013-07-01Request() *AccountAliyuncsComCreateAliyunAccount2013-07-01Request{
    return &AccountAliyuncsComCreateAliyunAccount2013-07-01Request{
        Params: model.NewParams(),
    }
}

func (r AccountAliyuncsComCreateAliyunAccount2013-07-01Request) GetApiMethodName() string {
    return "account.aliyuncs.com.CreateAliyunAccount.2013-07-01"
}

func (r AccountAliyuncsComCreateAliyunAccount2013-07-01Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


