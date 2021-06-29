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
type AlibabaWdkItemMerchantskuCreateRequest struct {
    model.Params
    // 新建商品参数，是个json字符串
    _params   string
}

// 初始化AlibabaWdkItemMerchantskuCreateRequest对象
func NewAlibabaWdkItemMerchantskuCreateRequest() *AlibabaWdkItemMerchantskuCreateRequest{
    return &AlibabaWdkItemMerchantskuCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMerchantskuCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.merchantsku.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMerchantskuCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Params Setter
// 新建商品参数，是个json字符串
func (r *AlibabaWdkItemMerchantskuCreateRequest) SetParams(_params string) error {
    r._params = _params
    r.Set("params", _params)
    return nil
}

// Params Getter
func (r AlibabaWdkItemMerchantskuCreateRequest) GetParams() string {
    return r._params
}
