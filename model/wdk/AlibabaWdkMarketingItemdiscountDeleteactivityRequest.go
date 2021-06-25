package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除商品特价活动 APIRequest
alibaba.wdk.marketing.itemdiscount.deleteactivity

删除商品特价活动
*/
type AlibabaWdkMarketingItemdiscountDeleteactivityRequest struct {
    model.Params

    // 需要删除的活动的信息
    param   *CommonActivityRequest 

}

func NewAlibabaWdkMarketingItemdiscountDeleteactivityRequest() *AlibabaWdkMarketingItemdiscountDeleteactivityRequest{
    return &AlibabaWdkMarketingItemdiscountDeleteactivityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItemdiscountDeleteactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itemdiscount.deleteactivity"
}

func (r AlibabaWdkMarketingItemdiscountDeleteactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItemdiscountDeleteactivityRequest) SetParam(param *CommonActivityRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingItemdiscountDeleteactivityRequest) GetParam() *CommonActivityRequest {
    return r.param
}

