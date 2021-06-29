package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新建商品草稿 API请求
alibaba.wdk.merchant.item.createdraft

新建商品草稿erp接口
*/
type AlibabaWdkMerchantItemCreatedraftRequest struct {
    model.Params
    // 商品信息json
    params   string
}

// 初始化AlibabaWdkMerchantItemCreatedraftRequest对象
func NewAlibabaWdkMerchantItemCreatedraftRequest() *AlibabaWdkMerchantItemCreatedraftRequest{
    return &AlibabaWdkMerchantItemCreatedraftRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantItemCreatedraftRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.item.createdraft"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantItemCreatedraftRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Params Setter
// 商品信息json
func (r *AlibabaWdkMerchantItemCreatedraftRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

// Params Getter
func (r AlibabaWdkMerchantItemCreatedraftRequest) GetParams() string {
    return r.params
}
