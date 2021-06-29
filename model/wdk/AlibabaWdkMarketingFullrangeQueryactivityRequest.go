package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全场活动查询活动 API请求
alibaba.wdk.marketing.fullrange.queryactivity

全场活动查询活动
*/
type AlibabaWdkMarketingFullrangeQueryactivityRequest struct {
    model.Params
    // 查询活动入参
    param0   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingFullrangeQueryactivityRequest对象
func NewAlibabaWdkMarketingFullrangeQueryactivityRequest() *AlibabaWdkMarketingFullrangeQueryactivityRequest{
    return &AlibabaWdkMarketingFullrangeQueryactivityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingFullrangeQueryactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.fullrange.queryactivity"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingFullrangeQueryactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 查询活动入参
func (r *AlibabaWdkMarketingFullrangeQueryactivityRequest) SetParam0(param0 *CommonActivityParam) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingFullrangeQueryactivityRequest) GetParam0() *CommonActivityParam {
    return r.param0
}
