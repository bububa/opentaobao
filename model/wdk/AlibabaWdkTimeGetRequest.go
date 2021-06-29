package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获得当前系统时间 API请求
alibaba.wdk.time.get

获得当前系统时间
*/
type AlibabaWdkTimeGetRequest struct {
    model.Params
}

// 初始化AlibabaWdkTimeGetRequest对象
func NewAlibabaWdkTimeGetRequest() *AlibabaWdkTimeGetRequest{
    return &AlibabaWdkTimeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkTimeGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.time.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkTimeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
