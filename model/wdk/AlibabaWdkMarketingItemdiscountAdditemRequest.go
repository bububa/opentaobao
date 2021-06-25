package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
报名特价商品 APIRequest
alibaba.wdk.marketing.itemdiscount.additem

在商品特价活动中报名特价商品
*/
type AlibabaWdkMarketingItemdiscountAdditemRequest struct {
    model.Params

    // 商品对象
    param0   *ItemDiscountSku 

    // 活动基本信息
    param1   *CommonActivityParam 

}

func NewAlibabaWdkMarketingItemdiscountAdditemRequest() *AlibabaWdkMarketingItemdiscountAdditemRequest{
    return &AlibabaWdkMarketingItemdiscountAdditemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItemdiscountAdditemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itemdiscount.additem"
}

func (r AlibabaWdkMarketingItemdiscountAdditemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItemdiscountAdditemRequest) SetParam0(param0 *ItemDiscountSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingItemdiscountAdditemRequest) GetParam0() *ItemDiscountSku {
    return r.param0
}

func (r *AlibabaWdkMarketingItemdiscountAdditemRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaWdkMarketingItemdiscountAdditemRequest) GetParam1() *CommonActivityParam {
    return r.param1
}

