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
type AlibabaWdkMarketingItempoolStairAdditemAPIRequest struct {
    model.Params
    // 系统自动生成
    _param0   *ItemPoolSku
    // 系统自动生成
    _param1   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItempoolStairAdditemAPIRequest对象
func NewAlibabaWdkMarketingItempoolStairAdditemRequest() *AlibabaWdkMarketingItempoolStairAdditemAPIRequest{
    return &AlibabaWdkMarketingItempoolStairAdditemAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolStairAdditemAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.stair.additem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolStairAdditemAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingItempoolStairAdditemAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItempoolStairAdditemAPIRequest) GetParam0() *ItemPoolSku {
    return r._param0
}
// Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingItempoolStairAdditemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingItempoolStairAdditemAPIRequest) GetParam1() *CommonActivityParam {
    return r._param1
}
