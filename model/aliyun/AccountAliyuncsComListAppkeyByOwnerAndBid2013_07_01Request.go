package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据渠道id和状态列出appkey API请求
account.aliyuncs.com.ListAppkeyByOwnerAndBid.2013-07-01

根据渠道id和状态列出appkey
*/
type AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01Request struct {
    model.Params
}

// 初始化AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01Request对象
func NewAccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01Request() *AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01Request{
    return &AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01Request{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01Request) GetApiMethodName() string {
    return "account.aliyuncs.com.ListAppkeyByOwnerAndBid.2013-07-01"
}

// IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
