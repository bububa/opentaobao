package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
移除买赠活动下的商品。【注意，此接口暂不支持并发！】 APIRequest
alibaba.wdk.marketing.itembuygift.removeitem

移除买赠活动下的商品。【注意，此接口暂不支持并发！】
*/
type AlibabaWdkMarketingItembuygiftRemoveitemRequest struct {
    model.Params

    // 商品对象
    param0   *ItemBuyGiftSku 

    // 活动基本信息
    param1   *CommonActivityParam 

}

func NewAlibabaWdkMarketingItembuygiftRemoveitemRequest() *AlibabaWdkMarketingItembuygiftRemoveitemRequest{
    return &AlibabaWdkMarketingItembuygiftRemoveitemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItembuygiftRemoveitemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itembuygift.removeitem"
}

func (r AlibabaWdkMarketingItembuygiftRemoveitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItembuygiftRemoveitemRequest) SetParam0(param0 *ItemBuyGiftSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingItembuygiftRemoveitemRequest) GetParam0() *ItemBuyGiftSku {
    return r.param0
}

func (r *AlibabaWdkMarketingItembuygiftRemoveitemRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaWdkMarketingItembuygiftRemoveitemRequest) GetParam1() *CommonActivityParam {
    return r.param1
}

