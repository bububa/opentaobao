package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除换购活动商品 APIRequest
alibaba.wdk.marketing.itempool.stair.removeitem

删除换购商品
*/
type AlibabaWdkMarketingItempoolStairRemoveitemRequest struct {
    model.Params

    // 商品sku信息
    param0   *ItemPoolSku 

    // 活动信息
    param1   *CommonActivityParam 

}

func NewAlibabaWdkMarketingItempoolStairRemoveitemRequest() *AlibabaWdkMarketingItempoolStairRemoveitemRequest{
    return &AlibabaWdkMarketingItempoolStairRemoveitemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItempoolStairRemoveitemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.stair.removeitem"
}

func (r AlibabaWdkMarketingItempoolStairRemoveitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItempoolStairRemoveitemRequest) SetParam0(param0 *ItemPoolSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingItempoolStairRemoveitemRequest) GetParam0() *ItemPoolSku {
    return r.param0
}

func (r *AlibabaWdkMarketingItempoolStairRemoveitemRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaWdkMarketingItempoolStairRemoveitemRequest) GetParam1() *CommonActivityParam {
    return r.param1
}

