package aliexpresssumaitong

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
平台固定参数获取 API请求
aliexpress.taxation.platform.open.get

Aliexpress开放平台固定参数获取
*/
type AliexpressTaxationPlatformOpenGetRequest struct {
    model.Params
}

// 初始化AliexpressTaxationPlatformOpenGetRequest对象
func NewAliexpressTaxationPlatformOpenGetRequest() *AliexpressTaxationPlatformOpenGetRequest{
    return &AliexpressTaxationPlatformOpenGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressTaxationPlatformOpenGetRequest) GetApiMethodName() string {
    return "aliexpress.taxation.platform.open.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressTaxationPlatformOpenGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
