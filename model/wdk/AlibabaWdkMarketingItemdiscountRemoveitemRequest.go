package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
移除报名的商品 APIRequest
alibaba.wdk.marketing.itemdiscount.removeitem

将报名特价活动的商品从特价活动中移除
*/
type AlibabaWdkMarketingItemdiscountRemoveitemRequest struct {
    model.Params

    // 商品对象
    param0   *ItemDiscountSku 

    // 活动基本信息
    param1   *CommonActivityParam 

}

func NewAlibabaWdkMarketingItemdiscountRemoveitemRequest() *AlibabaWdkMarketingItemdiscountRemoveitemRequest{
    return &AlibabaWdkMarketingItemdiscountRemoveitemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItemdiscountRemoveitemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itemdiscount.removeitem"
}

func (r AlibabaWdkMarketingItemdiscountRemoveitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItemdiscountRemoveitemRequest) SetParam0(param0 *ItemDiscountSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingItemdiscountRemoveitemRequest) GetParam0() *ItemDiscountSku {
    return r.param0
}

func (r *AlibabaWdkMarketingItemdiscountRemoveitemRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaWdkMarketingItemdiscountRemoveitemRequest) GetParam1() *CommonActivityParam {
    return r.param1
}

