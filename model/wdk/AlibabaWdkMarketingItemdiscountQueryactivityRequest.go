package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查找特价活动 APIRequest
alibaba.wdk.marketing.itemdiscount.queryactivity

查找特价活动
*/
type AlibabaWdkMarketingItemdiscountQueryactivityRequest struct {
    model.Params

    // 商品对象
    param0   *CommonActivityParam 

}

func NewAlibabaWdkMarketingItemdiscountQueryactivityRequest() *AlibabaWdkMarketingItemdiscountQueryactivityRequest{
    return &AlibabaWdkMarketingItemdiscountQueryactivityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItemdiscountQueryactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itemdiscount.queryactivity"
}

func (r AlibabaWdkMarketingItemdiscountQueryactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItemdiscountQueryactivityRequest) SetParam0(param0 *CommonActivityParam) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingItemdiscountQueryactivityRequest) GetParam0() *CommonActivityParam {
    return r.param0
}

