package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
换购商品查询 API请求
alibaba.wdk.marketing.itempool.stair.queryitem

换购商品查询
*/
type AlibabaWdkMarketingItempoolStairQueryitemRequest struct {
    model.Params
    // 换购商品查询参数
    param0   *ActivitySkuQuery
}

// 初始化AlibabaWdkMarketingItempoolStairQueryitemRequest对象
func NewAlibabaWdkMarketingItempoolStairQueryitemRequest() *AlibabaWdkMarketingItempoolStairQueryitemRequest{
    return &AlibabaWdkMarketingItempoolStairQueryitemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolStairQueryitemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.stair.queryitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolStairQueryitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 换购商品查询参数
func (r *AlibabaWdkMarketingItempoolStairQueryitemRequest) SetParam0(param0 *ActivitySkuQuery) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItempoolStairQueryitemRequest) GetParam0() *ActivitySkuQuery {
    return r.param0
}
