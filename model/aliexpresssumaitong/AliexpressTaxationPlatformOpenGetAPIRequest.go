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
type AliexpressTaxationPlatformOpenGetAPIRequest struct {
    model.Params
}

// 初始化AliexpressTaxationPlatformOpenGetAPIRequest对象
func NewAliexpressTaxationPlatformOpenGetRequest() *AliexpressTaxationPlatformOpenGetAPIRequest{
    return &AliexpressTaxationPlatformOpenGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressTaxationPlatformOpenGetAPIRequest) GetApiMethodName() string {
    return "aliexpress.taxation.platform.open.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressTaxationPlatformOpenGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
