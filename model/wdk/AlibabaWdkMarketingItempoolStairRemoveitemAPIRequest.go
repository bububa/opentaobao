package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除换购活动商品 API请求
alibaba.wdk.marketing.itempool.stair.removeitem

删除换购商品
*/
type AlibabaWdkMarketingItempoolStairRemoveitemAPIRequest struct {
    model.Params
    // 商品sku信息
    _param0   *ItemPoolSku
    // 活动信息
    _param1   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItempoolStairRemoveitemAPIRequest对象
func NewAlibabaWdkMarketingItempoolStairRemoveitemRequest() *AlibabaWdkMarketingItempoolStairRemoveitemAPIRequest{
    return &AlibabaWdkMarketingItempoolStairRemoveitemAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolStairRemoveitemAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.stair.removeitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolStairRemoveitemAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 商品sku信息
func (r *AlibabaWdkMarketingItempoolStairRemoveitemAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItempoolStairRemoveitemAPIRequest) GetParam0() *ItemPoolSku {
    return r._param0
}
// Param1 Setter
// 活动信息
func (r *AlibabaWdkMarketingItempoolStairRemoveitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingItempoolStairRemoveitemAPIRequest) GetParam1() *CommonActivityParam {
    return r._param1
}
