package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加商品池里面的商品 API请求
alibaba.wdk.marketing.itempool.additem

增加商品池里面的商品
*/
type AlibabaWdkMarketingItempoolAdditemAPIRequest struct {
    model.Params
    // 商品对象
    _param0   *ItemPoolSku
    // 活动基本信息
    _param1   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItempoolAdditemAPIRequest对象
func NewAlibabaWdkMarketingItempoolAdditemRequest() *AlibabaWdkMarketingItempoolAdditemAPIRequest{
    return &AlibabaWdkMarketingItempoolAdditemAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolAdditemAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.additem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolAdditemAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingItempoolAdditemAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItempoolAdditemAPIRequest) GetParam0() *ItemPoolSku {
    return r._param0
}
// Param1 Setter
// 活动基本信息
func (r *AlibabaWdkMarketingItempoolAdditemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingItempoolAdditemAPIRequest) GetParam1() *CommonActivityParam {
    return r._param1
}
