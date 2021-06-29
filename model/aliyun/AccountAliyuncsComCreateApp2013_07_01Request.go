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
type AccountAliyuncsComCreateApp2013_07_01Request struct {
    model.Params
}

// 初始化AccountAliyuncsComCreateApp2013_07_01Request对象
func NewAccountAliyuncsComCreateApp2013_07_01Request() *AccountAliyuncsComCreateApp2013_07_01Request{
    return &AccountAliyuncsComCreateApp2013_07_01Request{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComCreateApp2013_07_01Request) GetApiMethodName() string {
    return "account.aliyuncs.com.CreateApp.2013-07-01"
}

// IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComCreateApp2013_07_01Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
