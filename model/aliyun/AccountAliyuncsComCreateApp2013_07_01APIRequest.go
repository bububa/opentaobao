package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
给指定用户创建appkey API请求
account.aliyuncs.com.CreateApp.2013-07-01

为某个用户创建appkey
*/
type AccountAliyuncsComCreateApp2013_07_01APIRequest struct {
    model.Params
}

// 初始化AccountAliyuncsComCreateApp2013_07_01APIRequest对象
func NewAccountAliyuncsComCreateApp2013_07_01Request() *AccountAliyuncsComCreateApp2013_07_01APIRequest{
    return &AccountAliyuncsComCreateApp2013_07_01APIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComCreateApp2013_07_01APIRequest) GetApiMethodName() string {
    return "account.aliyuncs.com.CreateApp.2013-07-01"
}

// IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComCreateApp2013_07_01APIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
