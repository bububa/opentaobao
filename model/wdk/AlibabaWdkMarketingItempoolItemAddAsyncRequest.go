package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品池新增商品 API请求
alibaba.wdk.marketing.itempool.item.add.async

新分组模型下新增商品
*/
type AlibabaWdkMarketingItempoolItemAddAsyncRequest struct {
    model.Params
    // 阶梯商品信息
    param0   []ItemPoolSku
    // 系统自动生成
    param1   *CommonActivityParam
    // alibaba.wdk.marketing.version.generate接口生成
    version   int64
}

// 初始化AlibabaWdkMarketingItempoolItemAddAsyncRequest对象
func NewAlibabaWdkMarketingItempoolItemAddAsyncRequest() *AlibabaWdkMarketingItempoolItemAddAsyncRequest{
    return &AlibabaWdkMarketingItempoolItemAddAsyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolItemAddAsyncRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.item.add.async"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolItemAddAsyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 阶梯商品信息
func (r *AlibabaWdkMarketingItempoolItemAddAsyncRequest) SetParam0(param0 []ItemPoolSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItempoolItemAddAsyncRequest) GetParam0() []ItemPoolSku {
    return r.param0
}
// Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingItempoolItemAddAsyncRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingItempoolItemAddAsyncRequest) GetParam1() *CommonActivityParam {
    return r.param1
}
// Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaWdkMarketingItempoolItemAddAsyncRequest) SetVersion(version int64) error {
    r.version = version
    r.Set("version", version)
    return nil
}

// Version Getter
func (r AlibabaWdkMarketingItempoolItemAddAsyncRequest) GetVersion() int64 {
    return r.version
}
