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
type AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest struct {
    model.Params
    // 阶梯商品信息
    _param0   []ItemPoolSku
    // 系统自动生成
    _param1   *CommonActivityParam
    // alibaba.wdk.marketing.version.generate接口生成
    _version   int64
}

// 初始化AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest对象
func NewAlibabaWdkMarketingItempoolItemAddAsyncRequest() *AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest{
    return &AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.item.add.async"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 阶梯商品信息
func (r *AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) SetParam0(_param0 []ItemPoolSku) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) GetParam0() []ItemPoolSku {
    return r._param0
}
// Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) GetParam1() *CommonActivityParam {
    return r._param1
}
// Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) SetVersion(_version int64) error {
    r._version = _version
    r.Set("version", _version)
    return nil
}

// Version Getter
func (r AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) GetVersion() int64 {
    return r._version
}
