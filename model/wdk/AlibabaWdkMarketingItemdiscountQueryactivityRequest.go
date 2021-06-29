package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查找特价活动 API请求
alibaba.wdk.marketing.itemdiscount.queryactivity

查找特价活动
*/
type AlibabaWdkMarketingItemdiscountQueryactivityRequest struct {
    model.Params
    // 商品对象
    _param0   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItemdiscountQueryactivityRequest对象
func NewAlibabaWdkMarketingItemdiscountQueryactivityRequest() *AlibabaWdkMarketingItemdiscountQueryactivityRequest{
    return &AlibabaWdkMarketingItemdiscountQueryactivityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItemdiscountQueryactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itemdiscount.queryactivity"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItemdiscountQueryactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingItemdiscountQueryactivityRequest) SetParam0(_param0 *CommonActivityParam) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItemdiscountQueryactivityRequest) GetParam0() *CommonActivityParam {
    return r._param0
}
