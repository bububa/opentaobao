package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全场活动删除活动接口 APIRequest
alibaba.wdk.marketing.fullrange.deleteactivity

全场活动删除活动
*/
type AlibabaWdkMarketingFullrangeDeleteactivityRequest struct {
    model.Params

    // 需要删除的活动的信息
    param   *CommonActivityParam 

}

func NewAlibabaWdkMarketingFullrangeDeleteactivityRequest() *AlibabaWdkMarketingFullrangeDeleteactivityRequest{
    return &AlibabaWdkMarketingFullrangeDeleteactivityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingFullrangeDeleteactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.fullrange.deleteactivity"
}

func (r AlibabaWdkMarketingFullrangeDeleteactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingFullrangeDeleteactivityRequest) SetParam(param *CommonActivityParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingFullrangeDeleteactivityRequest) GetParam() *CommonActivityParam {
    return r.param
}

