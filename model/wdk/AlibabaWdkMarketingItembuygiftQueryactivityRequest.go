package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询买赠活动 APIRequest
alibaba.wdk.marketing.itembuygift.queryactivity

查询买赠活动
*/
type AlibabaWdkMarketingItembuygiftQueryactivityRequest struct {
    model.Params

    // 查询入参
    param   *CommonActivityParam 

}

func NewAlibabaWdkMarketingItembuygiftQueryactivityRequest() *AlibabaWdkMarketingItembuygiftQueryactivityRequest{
    return &AlibabaWdkMarketingItembuygiftQueryactivityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItembuygiftQueryactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itembuygift.queryactivity"
}

func (r AlibabaWdkMarketingItembuygiftQueryactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItembuygiftQueryactivityRequest) SetParam(param *CommonActivityParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingItembuygiftQueryactivityRequest) GetParam() *CommonActivityParam {
    return r.param
}

