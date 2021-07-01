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
type AlibabaWdkMarketingFullrangeRemoveitemAPIRequest struct {
    model.Params
    // 商品sku信息
    _param0   *ItemStairSku
    // 活动信息
    _param1   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingFullrangeRemoveitemAPIRequest对象
func NewAlibabaWdkMarketingFullrangeRemoveitemRequest() *AlibabaWdkMarketingFullrangeRemoveitemAPIRequest{
    return &AlibabaWdkMarketingFullrangeRemoveitemAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingFullrangeRemoveitemAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.fullrange.removeitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingFullrangeRemoveitemAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 商品sku信息
func (r *AlibabaWdkMarketingFullrangeRemoveitemAPIRequest) SetParam0(_param0 *ItemStairSku) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingFullrangeRemoveitemAPIRequest) GetParam0() *ItemStairSku {
    return r._param0
}
// Param1 Setter
// 活动信息
func (r *AlibabaWdkMarketingFullrangeRemoveitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingFullrangeRemoveitemAPIRequest) GetParam1() *CommonActivityParam {
    return r._param1
}
