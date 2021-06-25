package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全场活动查询换购品 APIRequest
alibaba.wdk.marketing.fullrange.queryitem

全场活动查询换购品
*/
type AlibabaWdkMarketingFullrangeQueryitemRequest struct {
    model.Params

    // 换购商品查询参数
    param0   *ActivitySkuQuery 

}

func NewAlibabaWdkMarketingFullrangeQueryitemRequest() *AlibabaWdkMarketingFullrangeQueryitemRequest{
    return &AlibabaWdkMarketingFullrangeQueryitemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingFullrangeQueryitemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.fullrange.queryitem"
}

func (r AlibabaWdkMarketingFullrangeQueryitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingFullrangeQueryitemRequest) SetParam0(param0 *ActivitySkuQuery) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingFullrangeQueryitemRequest) GetParam0() *ActivitySkuQuery {
    return r.param0
}

