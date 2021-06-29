package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除商品池活动 API请求
alibaba.wdk.marketing.itempool.deleteactivity

删除商品池活动
*/
type AlibabaWdkMarketingItempoolDeleteactivityRequest struct {
    model.Params
    // 需要删除的活动的信息
    param   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItempoolDeleteactivityRequest对象
func NewAlibabaWdkMarketingItempoolDeleteactivityRequest() *AlibabaWdkMarketingItempoolDeleteactivityRequest{
    return &AlibabaWdkMarketingItempoolDeleteactivityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolDeleteactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.deleteactivity"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolDeleteactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 需要删除的活动的信息
func (r *AlibabaWdkMarketingItempoolDeleteactivityRequest) SetParam(param *CommonActivityParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingItempoolDeleteactivityRequest) GetParam() *CommonActivityParam {
    return r.param
}
