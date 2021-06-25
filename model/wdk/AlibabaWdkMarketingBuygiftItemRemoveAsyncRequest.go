package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除买赠商品 APIRequest
alibaba.wdk.marketing.buygift.item.remove.async

批量删除买赠商品
*/
type AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest struct {
    model.Params

    // sku信息
    param0   []ItemBuyGiftSku 

    // 系统自动生成
    param1   *CommonActivityParam 

    // alibaba.wdk.marketing.version.generate接口生成
    version   int64 

}

func NewAlibabaWdkMarketingBuygiftItemRemoveAsyncRequest() *AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest{
    return &AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.buygift.item.remove.async"
}

func (r AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest) SetParam0(param0 []ItemBuyGiftSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest) GetParam0() []ItemBuyGiftSku {
    return r.param0
}

func (r *AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest) GetParam1() *CommonActivityParam {
    return r.param1
}

func (r *AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest) SetVersion(version int64) error {
    r.version = version
    r.Set("version", version)
    return nil
}

func (r AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest) GetVersion() int64 {
    return r.version
}

