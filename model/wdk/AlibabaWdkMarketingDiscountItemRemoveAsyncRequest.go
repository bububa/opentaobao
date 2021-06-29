package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
特价批量移除商品 API请求
alibaba.wdk.marketing.discount.item.remove.async

特价批量移除商品
*/
type AlibabaWdkMarketingDiscountItemRemoveAsyncRequest struct {
    model.Params
    // sku信息
    _param0   []ItemDiscountSku
    // 系统自动生成
    _param1   *CommonActivityParam
    // alibaba.wdk.marketing.version.generate接口生成
    _version   int64
}

// 初始化AlibabaWdkMarketingDiscountItemRemoveAsyncRequest对象
func NewAlibabaWdkMarketingDiscountItemRemoveAsyncRequest() *AlibabaWdkMarketingDiscountItemRemoveAsyncRequest{
    return &AlibabaWdkMarketingDiscountItemRemoveAsyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingDiscountItemRemoveAsyncRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.discount.item.remove.async"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingDiscountItemRemoveAsyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// sku信息
func (r *AlibabaWdkMarketingDiscountItemRemoveAsyncRequest) SetParam0(_param0 []ItemDiscountSku) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingDiscountItemRemoveAsyncRequest) GetParam0() []ItemDiscountSku {
    return r._param0
}
// Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingDiscountItemRemoveAsyncRequest) SetParam1(_param1 *CommonActivityParam) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingDiscountItemRemoveAsyncRequest) GetParam1() *CommonActivityParam {
    return r._param1
}
// Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaWdkMarketingDiscountItemRemoveAsyncRequest) SetVersion(_version int64) error {
    r._version = _version
    r.Set("version", _version)
    return nil
}

// Version Getter
func (r AlibabaWdkMarketingDiscountItemRemoveAsyncRequest) GetVersion() int64 {
    return r._version
}
