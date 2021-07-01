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
type AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest struct {
    model.Params
    // sku信息
    _param0   []ItemPoolSku
    // 活动信息
    _param1   *CommonActivityParam
    // alibaba.wdk.marketing.version.generate接口生成
    _version   int64
}

// 初始化AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest对象
func NewAlibabaWdkMarketingItempoolItemRemoveAsyncRequest() *AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest{
    return &AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.item.remove.async"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// sku信息
func (r *AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) SetParam0(_param0 []ItemPoolSku) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) GetParam0() []ItemPoolSku {
    return r._param0
}
// Param1 Setter
// 活动信息
func (r *AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) GetParam1() *CommonActivityParam {
    return r._param1
}
// Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) SetVersion(_version int64) error {
    r._version = _version
    r.Set("version", _version)
    return nil
}

// Version Getter
func (r AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) GetVersion() int64 {
    return r._version
}
