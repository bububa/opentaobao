package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除买赠活动 APIRequest
alibaba.wdk.marketing.itembuygift.deleteactivity

删除买赠活动
*/
type AlibabaWdkMarketingItembuygiftDeleteactivityRequest struct {
    model.Params

    // 要删除的活动信息
    param   *CommonActivityParam 

}

func NewAlibabaWdkMarketingItembuygiftDeleteactivityRequest() *AlibabaWdkMarketingItembuygiftDeleteactivityRequest{
    return &AlibabaWdkMarketingItembuygiftDeleteactivityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItembuygiftDeleteactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itembuygift.deleteactivity"
}

func (r AlibabaWdkMarketingItembuygiftDeleteactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItembuygiftDeleteactivityRequest) SetParam(param *CommonActivityParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingItembuygiftDeleteactivityRequest) GetParam() *CommonActivityParam {
    return r.param
}

