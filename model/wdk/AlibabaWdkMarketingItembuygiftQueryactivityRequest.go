package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询买赠活动 API请求
alibaba.wdk.marketing.itembuygift.queryactivity

查询买赠活动
*/
type AlibabaWdkMarketingItembuygiftQueryactivityRequest struct {
    model.Params
    // 查询入参
    _param   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItembuygiftQueryactivityRequest对象
func NewAlibabaWdkMarketingItembuygiftQueryactivityRequest() *AlibabaWdkMarketingItembuygiftQueryactivityRequest{
    return &AlibabaWdkMarketingItembuygiftQueryactivityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItembuygiftQueryactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itembuygift.queryactivity"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItembuygiftQueryactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 查询入参
func (r *AlibabaWdkMarketingItembuygiftQueryactivityRequest) SetParam(_param *CommonActivityParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingItembuygiftQueryactivityRequest) GetParam() *CommonActivityParam {
    return r._param
}
