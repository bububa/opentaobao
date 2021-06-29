package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全场活动查询换购品 API请求
alibaba.wdk.marketing.fullrange.queryitem

全场活动查询换购品
*/
type AlibabaWdkMarketingFullrangeQueryitemRequest struct {
    model.Params
    // 换购商品查询参数
    _param0   *ActivitySkuQuery
}

// 初始化AlibabaWdkMarketingFullrangeQueryitemRequest对象
func NewAlibabaWdkMarketingFullrangeQueryitemRequest() *AlibabaWdkMarketingFullrangeQueryitemRequest{
    return &AlibabaWdkMarketingFullrangeQueryitemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingFullrangeQueryitemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.fullrange.queryitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingFullrangeQueryitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 换购商品查询参数
func (r *AlibabaWdkMarketingFullrangeQueryitemRequest) SetParam0(_param0 *ActivitySkuQuery) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingFullrangeQueryitemRequest) GetParam0() *ActivitySkuQuery {
    return r._param0
}
