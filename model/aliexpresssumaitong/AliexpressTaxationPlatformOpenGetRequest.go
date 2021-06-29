package aliexpresssumaitong

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
平台固定参数获取 APIRequest
aliexpress.taxation.platform.open.get

Aliexpress开放平台固定参数获取
*/
type AliexpressTaxationPlatformOpenGetRequest struct {
    model.Params

}

func NewAliexpressTaxationPlatformOpenGetRequest() *AliexpressTaxationPlatformOpenGetRequest{
    return &AliexpressTaxationPlatformOpenGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressTaxationPlatformOpenGetRequest) GetApiMethodName() string {
    return "aliexpress.taxation.platform.open.get"
}

func (r AliexpressTaxationPlatformOpenGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


