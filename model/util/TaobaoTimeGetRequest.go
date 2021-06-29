package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取淘宝系统当前时间 API请求
taobao.time.get

获取淘宝系统当前时间
*/
type TaobaoTimeGetRequest struct {
    model.Params
}

// 初始化TaobaoTimeGetRequest对象
func NewTaobaoTimeGetRequest() *TaobaoTimeGetRequest{
    return &TaobaoTimeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTimeGetRequest) GetApiMethodName() string {
    return "taobao.time.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTimeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
