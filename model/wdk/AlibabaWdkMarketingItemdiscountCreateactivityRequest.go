package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建商品特价活动 APIRequest
alibaba.wdk.marketing.itemdiscount.createactivity

创建商品特价活动
*/
type AlibabaWdkMarketingItemdiscountCreateactivityRequest struct {
    model.Params

    // 创建活动请求入参
    param   *ItemDiscountActivityRequest 

}

func NewAlibabaWdkMarketingItemdiscountCreateactivityRequest() *AlibabaWdkMarketingItemdiscountCreateactivityRequest{
    return &AlibabaWdkMarketingItemdiscountCreateactivityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItemdiscountCreateactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itemdiscount.createactivity"
}

func (r AlibabaWdkMarketingItemdiscountCreateactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItemdiscountCreateactivityRequest) SetParam(param *ItemDiscountActivityRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingItemdiscountCreateactivityRequest) GetParam() *ItemDiscountActivityRequest {
    return r.param
}

