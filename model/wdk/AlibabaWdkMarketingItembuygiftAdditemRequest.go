package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加买赠活动商品。【注意，此接口暂不支持并发！】 APIRequest
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

func NewAlibabaWdkMarketingItembuygiftAdditemRequest() *AlibabaWdkMarketingItembuygiftAdditemRequest{
    return &AlibabaWdkMarketingItembuygiftAdditemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItembuygiftAdditemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itembuygift.additem"
}

func (r AlibabaWdkMarketingItembuygiftAdditemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItembuygiftAdditemRequest) SetParam0(param0 *ItemBuyGiftSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingItembuygiftAdditemRequest) GetParam0() *ItemBuyGiftSku {
    return r.param0
}

func (r *AlibabaWdkMarketingItembuygiftAdditemRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaWdkMarketingItembuygiftAdditemRequest) GetParam1() *CommonActivityParam {
    return r.param1
}

