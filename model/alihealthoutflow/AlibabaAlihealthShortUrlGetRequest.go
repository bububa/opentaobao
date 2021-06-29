package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付宝短链跳转三方h5通用接口 API请求
alibaba.alihealth.short.url.get

支付宝短链跳转三方h5通用接口
*/
type AlibabaAlihealthShortUrlGetRequest struct {
    model.Params
    // 三方h5
    _url   string
    // 参数替换列表
    _params   []string
}

// 初始化AlibabaAlihealthShortUrlGetRequest对象
func NewAlibabaAlihealthShortUrlGetRequest() *AlibabaAlihealthShortUrlGetRequest{
    return &AlibabaAlihealthShortUrlGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthShortUrlGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.short.url.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthShortUrlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Url Setter
// 三方h5
func (r *AlibabaAlihealthShortUrlGetRequest) SetUrl(_url string) error {
    r._url = _url
    r.Set("url", _url)
    return nil
}

// Url Getter
func (r AlibabaAlihealthShortUrlGetRequest) GetUrl() string {
    return r._url
}
// Params Setter
// 参数替换列表
func (r *AlibabaAlihealthShortUrlGetRequest) SetParams(_params []string) error {
    r._params = _params
    r.Set("params", _params)
    return nil
}

// Params Getter
func (r AlibabaAlihealthShortUrlGetRequest) GetParams() []string {
    return r._params
}
