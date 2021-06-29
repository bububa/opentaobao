package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全场活动删除活动接口 API请求
alibaba.wdk.marketing.fullrange.deleteactivity

全场活动删除活动
*/
type AlibabaWdkMarketingFullrangeDeleteactivityRequest struct {
    model.Params
    // 需要删除的活动的信息
    param   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingFullrangeDeleteactivityRequest对象
func NewAlibabaWdkMarketingFullrangeDeleteactivityRequest() *AlibabaWdkMarketingFullrangeDeleteactivityRequest{
    return &AlibabaWdkMarketingFullrangeDeleteactivityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingFullrangeDeleteactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.fullrange.deleteactivity"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingFullrangeDeleteactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 需要删除的活动的信息
func (r *AlibabaWdkMarketingFullrangeDeleteactivityRequest) SetParam(param *CommonActivityParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingFullrangeDeleteactivityRequest) GetParam() *CommonActivityParam {
    return r.param
}
