package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加买赠活动商品。【注意，此接口暂不支持并发！】 API请求
alibaba.wdk.marketing.itembuygift.additem

增加买赠活动商品。【注意，此接口暂不支持并发！】
*/
type AlibabaWdkMarketingItembuygiftAdditemRequest struct {
    model.Params
    // 商品对象
    param0   *ItemBuyGiftSku
    // 活动基本信息
    param1   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItembuygiftAdditemRequest对象
func NewAlibabaWdkMarketingItembuygiftAdditemRequest() *AlibabaWdkMarketingItembuygiftAdditemRequest{
    return &AlibabaWdkMarketingItembuygiftAdditemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItembuygiftAdditemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itembuygift.additem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItembuygiftAdditemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingItembuygiftAdditemRequest) SetParam0(param0 *ItemBuyGiftSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItembuygiftAdditemRequest) GetParam0() *ItemBuyGiftSku {
    return r.param0
}
// Param1 Setter
// 活动基本信息
func (r *AlibabaWdkMarketingItembuygiftAdditemRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingItembuygiftAdditemRequest) GetParam1() *CommonActivityParam {
    return r.param1
}
