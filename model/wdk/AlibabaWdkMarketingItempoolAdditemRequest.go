package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加商品池里面的商品 APIRequest
alibaba.wdk.marketing.itempool.additem

增加商品池里面的商品
*/
type AlibabaWdkMarketingItempoolAdditemRequest struct {
    model.Params

    // 商品对象
    param0   *ItemPoolSku 

    // 活动基本信息
    param1   *CommonActivityParam 

}

func NewAlibabaWdkMarketingItempoolAdditemRequest() *AlibabaWdkMarketingItempoolAdditemRequest{
    return &AlibabaWdkMarketingItempoolAdditemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItempoolAdditemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.additem"
}

func (r AlibabaWdkMarketingItempoolAdditemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItempoolAdditemRequest) SetParam0(param0 *ItemPoolSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingItempoolAdditemRequest) GetParam0() *ItemPoolSku {
    return r.param0
}

func (r *AlibabaWdkMarketingItempoolAdditemRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaWdkMarketingItempoolAdditemRequest) GetParam1() *CommonActivityParam {
    return r.param1
}

