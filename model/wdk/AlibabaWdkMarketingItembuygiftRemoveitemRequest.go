package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
移除买赠活动下的商品。【注意，此接口暂不支持并发！】 API请求
alibaba.wdk.marketing.itembuygift.removeitem

移除买赠活动下的商品。【注意，此接口暂不支持并发！】
*/
type AlibabaWdkMarketingItembuygiftRemoveitemRequest struct {
    model.Params
    // 商品对象
    _param0   *ItemBuyGiftSku
    // 活动基本信息
    _param1   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItembuygiftRemoveitemRequest对象
func NewAlibabaWdkMarketingItembuygiftRemoveitemRequest() *AlibabaWdkMarketingItembuygiftRemoveitemRequest{
    return &AlibabaWdkMarketingItembuygiftRemoveitemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItembuygiftRemoveitemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itembuygift.removeitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItembuygiftRemoveitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingItembuygiftRemoveitemRequest) SetParam0(_param0 *ItemBuyGiftSku) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItembuygiftRemoveitemRequest) GetParam0() *ItemBuyGiftSku {
    return r._param0
}
// Param1 Setter
// 活动基本信息
func (r *AlibabaWdkMarketingItembuygiftRemoveitemRequest) SetParam1(_param1 *CommonActivityParam) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingItembuygiftRemoveitemRequest) GetParam1() *CommonActivityParam {
    return r._param1
}
