package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除商品池活动 APIRequest
alibaba.wdk.marketing.itempool.deleteactivity

删除商品池活动
*/
type AlibabaWdkMarketingItempoolDeleteactivityRequest struct {
    model.Params

    // 需要删除的活动的信息
    param   *CommonActivityParam 

}

func NewAlibabaWdkMarketingItempoolDeleteactivityRequest() *AlibabaWdkMarketingItempoolDeleteactivityRequest{
    return &AlibabaWdkMarketingItempoolDeleteactivityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItempoolDeleteactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.deleteactivity"
}

func (r AlibabaWdkMarketingItempoolDeleteactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItempoolDeleteactivityRequest) SetParam(param *CommonActivityParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingItempoolDeleteactivityRequest) GetParam() *CommonActivityParam {
    return r.param
}

