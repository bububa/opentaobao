package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付宝短链跳转三方h5通用接口 APIRequest
alibaba.alihealth.short.url.get

支付宝短链跳转三方h5通用接口
*/
type AlibabaAlihealthShortUrlGetRequest struct {
    model.Params

    // 三方h5
    url   string 

    // 参数替换列表
    params   []string 

}

func NewAlibabaAlihealthShortUrlGetRequest() *AlibabaAlihealthShortUrlGetRequest{
    return &AlibabaAlihealthShortUrlGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthShortUrlGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.short.url.get"
}

func (r AlibabaAlihealthShortUrlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthShortUrlGetRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

func (r AlibabaAlihealthShortUrlGetRequest) GetUrl() string {
    return r.url
}

func (r *AlibabaAlihealthShortUrlGetRequest) SetParams(params []string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

func (r AlibabaAlihealthShortUrlGetRequest) GetParams() []string {
    return r.params
}

