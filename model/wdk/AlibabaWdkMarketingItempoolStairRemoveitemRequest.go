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
type AlibabaWdkMarketingItempoolStairRemoveitemRequest struct {
    model.Params
    // 商品sku信息
    param0   *ItemPoolSku
    // 活动信息
    param1   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItempoolStairRemoveitemRequest对象
func NewAlibabaWdkMarketingItempoolStairRemoveitemRequest() *AlibabaWdkMarketingItempoolStairRemoveitemRequest{
    return &AlibabaWdkMarketingItempoolStairRemoveitemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolStairRemoveitemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.stair.removeitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolStairRemoveitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 商品sku信息
func (r *AlibabaWdkMarketingItempoolStairRemoveitemRequest) SetParam0(param0 *ItemPoolSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItempoolStairRemoveitemRequest) GetParam0() *ItemPoolSku {
    return r.param0
}
// Param1 Setter
// 活动信息
func (r *AlibabaWdkMarketingItempoolStairRemoveitemRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingItempoolStairRemoveitemRequest) GetParam1() *CommonActivityParam {
    return r.param1
}
