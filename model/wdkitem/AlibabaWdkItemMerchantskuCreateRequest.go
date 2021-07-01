package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品信息新建 API请求
alibaba.wdk.item.merchantsku.create

商家商品信息新建
*/
type AlibabaWdkItemMerchantskuCreateAPIRequest struct {
    model.Params
    // 新建商品参数，是个json字符串
    _params   string
}

// 初始化AlibabaWdkItemMerchantskuCreateAPIRequest对象
func NewAlibabaWdkItemMerchantskuCreateRequest() *AlibabaWdkItemMerchantskuCreateAPIRequest{
    return &AlibabaWdkItemMerchantskuCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMerchantskuCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.merchantsku.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMerchantskuCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Params Setter
// 新建商品参数，是个json字符串
func (r *AlibabaWdkItemMerchantskuCreateAPIRequest) SetParams(_params string) error {
    r._params = _params
    r.Set("params", _params)
    return nil
}

// Params Getter
func (r AlibabaWdkItemMerchantskuCreateAPIRequest) GetParams() string {
    return r._params
}
