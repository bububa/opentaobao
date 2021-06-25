package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获得当前系统时间 APIRequest
alibaba.wdk.time.get

获得当前系统时间
*/
type AlibabaWdkTimeGetRequest struct {
    model.Params

}

func NewAlibabaWdkTimeGetRequest() *AlibabaWdkTimeGetRequest{
    return &AlibabaWdkTimeGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkTimeGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.time.get"
}

func (r AlibabaWdkTimeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


