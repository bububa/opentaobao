package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量发布买赠商品 APIRequest
alibaba.wdk.marketing.buygift.item.add.async

批量发布买赠商品
*/
type AlibabaWdkMarketingBuygiftItemAddAsyncRequest struct {
    model.Params

    // sku信息
    param0   []ItemBuyGiftSku 

    // 系统自动生成
    param1   *CommonActivityParam 

    // alibaba.wdk.marketing.version.generate接口生成
    version   int64 

}

func NewAlibabaWdkMarketingBuygiftItemAddAsyncRequest() *AlibabaWdkMarketingBuygiftItemAddAsyncRequest{
    return &AlibabaWdkMarketingBuygiftItemAddAsyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingBuygiftItemAddAsyncRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.buygift.item.add.async"
}

func (r AlibabaWdkMarketingBuygiftItemAddAsyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingBuygiftItemAddAsyncRequest) SetParam0(param0 []ItemBuyGiftSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingBuygiftItemAddAsyncRequest) GetParam0() []ItemBuyGiftSku {
    return r.param0
}

func (r *AlibabaWdkMarketingBuygiftItemAddAsyncRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaWdkMarketingBuygiftItemAddAsyncRequest) GetParam1() *CommonActivityParam {
    return r.param1
}

func (r *AlibabaWdkMarketingBuygiftItemAddAsyncRequest) SetVersion(version int64) error {
    r.version = version
    r.Set("version", version)
    return nil
}

func (r AlibabaWdkMarketingBuygiftItemAddAsyncRequest) GetVersion() int64 {
    return r.version
}

