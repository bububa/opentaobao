package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全场活动删除购品 API请求
alibaba.wdk.marketing.fullrange.removeitem

删除换购商品
*/
type AlibabaWdkMarketingFullrangeRemoveitemRequest struct {
    model.Params
    // 商品sku信息
    param0   *ItemStairSku
    // 活动信息
    param1   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingFullrangeRemoveitemRequest对象
func NewAlibabaWdkMarketingFullrangeRemoveitemRequest() *AlibabaWdkMarketingFullrangeRemoveitemRequest{
    return &AlibabaWdkMarketingFullrangeRemoveitemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingFullrangeRemoveitemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.fullrange.removeitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingFullrangeRemoveitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 商品sku信息
func (r *AlibabaWdkMarketingFullrangeRemoveitemRequest) SetParam0(param0 *ItemStairSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingFullrangeRemoveitemRequest) GetParam0() *ItemStairSku {
    return r.param0
}
// Param1 Setter
// 活动信息
func (r *AlibabaWdkMarketingFullrangeRemoveitemRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingFullrangeRemoveitemRequest) GetParam1() *CommonActivityParam {
    return r.param1
}
