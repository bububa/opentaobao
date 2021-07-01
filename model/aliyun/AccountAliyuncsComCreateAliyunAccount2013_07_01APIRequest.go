package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建阿里云账号 API请求
account.aliyuncs.com.CreateAliyunAccount.2013-07-01

根据给定的阿里云账号，密码以及手机号创建阿里云账号
*/
type AccountAliyuncsComCreateAliyunAccount2013_07_01APIRequest struct {
    model.Params
}

// 初始化AccountAliyuncsComCreateAliyunAccount2013_07_01APIRequest对象
func NewAccountAliyuncsComCreateAliyunAccount2013_07_01Request() *AccountAliyuncsComCreateAliyunAccount2013_07_01APIRequest{
    return &AccountAliyuncsComCreateAliyunAccount2013_07_01APIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComCreateAliyunAccount2013_07_01APIRequest) GetApiMethodName() string {
    return "account.aliyuncs.com.CreateAliyunAccount.2013-07-01"
}

// IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComCreateAliyunAccount2013_07_01APIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
