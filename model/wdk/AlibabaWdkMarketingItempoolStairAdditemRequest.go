package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品池阶梯商品添加 API请求
alibaba.wdk.marketing.itempool.stair.additem

添加商品池阶梯商品
*/
type AlibabaWdkMarketingItempoolStairAdditemRequest struct {
    model.Params
    // 系统自动生成
    param0   *ItemPoolSku
    // 系统自动生成
    param1   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItempoolStairAdditemRequest对象
func NewAlibabaWdkMarketingItempoolStairAdditemRequest() *AlibabaWdkMarketingItempoolStairAdditemRequest{
    return &AlibabaWdkMarketingItempoolStairAdditemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolStairAdditemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.stair.additem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolStairAdditemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingItempoolStairAdditemRequest) SetParam0(param0 *ItemPoolSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItempoolStairAdditemRequest) GetParam0() *ItemPoolSku {
    return r.param0
}
// Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingItempoolStairAdditemRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingItempoolStairAdditemRequest) GetParam1() *CommonActivityParam {
    return r.param1
}
