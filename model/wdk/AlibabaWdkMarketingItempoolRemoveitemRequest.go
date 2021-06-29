package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
移除商品池里面的商品 API请求
alibaba.wdk.marketing.itempool.removeitem

移除商品池里面的商品
*/
type AlibabaWdkMarketingItempoolRemoveitemRequest struct {
    model.Params
    // 商品对象
    _param0   *ItemPoolSku
    // 活动基本信息
    _param1   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItempoolRemoveitemRequest对象
func NewAlibabaWdkMarketingItempoolRemoveitemRequest() *AlibabaWdkMarketingItempoolRemoveitemRequest{
    return &AlibabaWdkMarketingItempoolRemoveitemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolRemoveitemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.removeitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolRemoveitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingItempoolRemoveitemRequest) SetParam0(_param0 *ItemPoolSku) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItempoolRemoveitemRequest) GetParam0() *ItemPoolSku {
    return r._param0
}
// Param1 Setter
// 活动基本信息
func (r *AlibabaWdkMarketingItempoolRemoveitemRequest) SetParam1(_param1 *CommonActivityParam) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingItempoolRemoveitemRequest) GetParam1() *CommonActivityParam {
    return r._param1
}
