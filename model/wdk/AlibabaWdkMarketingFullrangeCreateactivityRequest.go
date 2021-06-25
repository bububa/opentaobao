package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建全场活动 APIRequest
alibaba.wdk.marketing.fullrange.createactivity

创建全场活动
*/
type AlibabaWdkMarketingFullrangeCreateactivityRequest struct {
    model.Params

    // 创建活动请求入参
    param   *FullRangeActivity 

}

func NewAlibabaWdkMarketingFullrangeCreateactivityRequest() *AlibabaWdkMarketingFullrangeCreateactivityRequest{
    return &AlibabaWdkMarketingFullrangeCreateactivityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingFullrangeCreateactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.fullrange.createactivity"
}

func (r AlibabaWdkMarketingFullrangeCreateactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingFullrangeCreateactivityRequest) SetParam(param *FullRangeActivity) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingFullrangeCreateactivityRequest) GetParam() *FullRangeActivity {
    return r.param
}

