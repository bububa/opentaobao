package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品池删除商品 API请求
alibaba.wdk.marketing.itempool.item.remove.async

新模型下删除商品
*/
type AlibabaWdkMarketingItempoolItemRemoveAsyncRequest struct {
    model.Params
    // sku信息
    param0   []ItemPoolSku
    // 活动信息
    param1   *CommonActivityParam
    // alibaba.wdk.marketing.version.generate接口生成
    version   int64
}

// 初始化AlibabaWdkMarketingItempoolItemRemoveAsyncRequest对象
func NewAlibabaWdkMarketingItempoolItemRemoveAsyncRequest() *AlibabaWdkMarketingItempoolItemRemoveAsyncRequest{
    return &AlibabaWdkMarketingItempoolItemRemoveAsyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolItemRemoveAsyncRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.item.remove.async"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolItemRemoveAsyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// sku信息
func (r *AlibabaWdkMarketingItempoolItemRemoveAsyncRequest) SetParam0(param0 []ItemPoolSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItempoolItemRemoveAsyncRequest) GetParam0() []ItemPoolSku {
    return r.param0
}
// Param1 Setter
// 活动信息
func (r *AlibabaWdkMarketingItempoolItemRemoveAsyncRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingItempoolItemRemoveAsyncRequest) GetParam1() *CommonActivityParam {
    return r.param1
}
// Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaWdkMarketingItempoolItemRemoveAsyncRequest) SetVersion(version int64) error {
    r.version = version
    r.Set("version", version)
    return nil
}

// Version Getter
func (r AlibabaWdkMarketingItempoolItemRemoveAsyncRequest) GetVersion() int64 {
    return r.version
}
