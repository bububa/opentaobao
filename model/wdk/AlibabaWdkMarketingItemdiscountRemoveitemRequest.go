package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
移除报名的商品 API请求
alibaba.wdk.marketing.itemdiscount.removeitem

将报名特价活动的商品从特价活动中移除
*/
type AlibabaWdkMarketingItemdiscountRemoveitemRequest struct {
    model.Params
    // 商品对象
    _param0   *ItemDiscountSku
    // 活动基本信息
    _param1   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItemdiscountRemoveitemRequest对象
func NewAlibabaWdkMarketingItemdiscountRemoveitemRequest() *AlibabaWdkMarketingItemdiscountRemoveitemRequest{
    return &AlibabaWdkMarketingItemdiscountRemoveitemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItemdiscountRemoveitemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itemdiscount.removeitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItemdiscountRemoveitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingItemdiscountRemoveitemRequest) SetParam0(_param0 *ItemDiscountSku) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItemdiscountRemoveitemRequest) GetParam0() *ItemDiscountSku {
    return r._param0
}
// Param1 Setter
// 活动基本信息
func (r *AlibabaWdkMarketingItemdiscountRemoveitemRequest) SetParam1(_param1 *CommonActivityParam) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingItemdiscountRemoveitemRequest) GetParam1() *CommonActivityParam {
    return r._param1
}
