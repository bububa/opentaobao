package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除买赠商品 API请求
alibaba.wdk.marketing.buygift.item.remove.async

批量删除买赠商品
*/
type AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest struct {
    model.Params
    // sku信息
    _param0   []ItemBuyGiftSku
    // 系统自动生成
    _param1   *CommonActivityParam
    // alibaba.wdk.marketing.version.generate接口生成
    _version   int64
}

// 初始化AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest对象
func NewAlibabaWdkMarketingBuygiftItemRemoveAsyncRequest() *AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest{
    return &AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.buygift.item.remove.async"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// sku信息
func (r *AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest) SetParam0(_param0 []ItemBuyGiftSku) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest) GetParam0() []ItemBuyGiftSku {
    return r._param0
}
// Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest) SetParam1(_param1 *CommonActivityParam) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest) GetParam1() *CommonActivityParam {
    return r._param1
}
// Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest) SetVersion(_version int64) error {
    r._version = _version
    r.Set("version", _version)
    return nil
}

// Version Getter
func (r AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest) GetVersion() int64 {
    return r._version
}
