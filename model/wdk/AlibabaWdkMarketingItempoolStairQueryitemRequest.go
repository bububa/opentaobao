package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
换购商品查询 APIRequest
alibaba.wdk.marketing.itempool.stair.queryitem

换购商品查询
*/
type AlibabaWdkMarketingItempoolStairQueryitemRequest struct {
    model.Params

    // 换购商品查询参数
    param0   *ActivitySkuQuery 

}

func NewAlibabaWdkMarketingItempoolStairQueryitemRequest() *AlibabaWdkMarketingItempoolStairQueryitemRequest{
    return &AlibabaWdkMarketingItempoolStairQueryitemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItempoolStairQueryitemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.stair.queryitem"
}

func (r AlibabaWdkMarketingItempoolStairQueryitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItempoolStairQueryitemRequest) SetParam0(param0 *ActivitySkuQuery) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingItempoolStairQueryitemRequest) GetParam0() *ActivitySkuQuery {
    return r.param0
}

