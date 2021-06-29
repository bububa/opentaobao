package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查找商品池活动 API请求
alibaba.wdk.marketing.itempool.queryactivity

查找商品池活动
*/
type AlibabaWdkMarketingItempoolQueryactivityRequest struct {
    model.Params
    // 查询商品池活动入参
    param0   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItempoolQueryactivityRequest对象
func NewAlibabaWdkMarketingItempoolQueryactivityRequest() *AlibabaWdkMarketingItempoolQueryactivityRequest{
    return &AlibabaWdkMarketingItempoolQueryactivityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolQueryactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.queryactivity"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolQueryactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 查询商品池活动入参
func (r *AlibabaWdkMarketingItempoolQueryactivityRequest) SetParam0(param0 *CommonActivityParam) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItempoolQueryactivityRequest) GetParam0() *CommonActivityParam {
    return r.param0
}
