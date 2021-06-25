package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
特价批量移除商品 APIRequest
alibaba.wdk.marketing.discount.item.remove.async

特价批量移除商品
*/
type AlibabaWdkMarketingDiscountItemRemoveAsyncRequest struct {
    model.Params

    // sku信息
    param0   []ItemDiscountSku 

    // 系统自动生成
    param1   *CommonActivityParam 

    // alibaba.wdk.marketing.version.generate接口生成
    version   int64 

}

func NewAlibabaWdkMarketingDiscountItemRemoveAsyncRequest() *AlibabaWdkMarketingDiscountItemRemoveAsyncRequest{
    return &AlibabaWdkMarketingDiscountItemRemoveAsyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingDiscountItemRemoveAsyncRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.discount.item.remove.async"
}

func (r AlibabaWdkMarketingDiscountItemRemoveAsyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingDiscountItemRemoveAsyncRequest) SetParam0(param0 []ItemDiscountSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingDiscountItemRemoveAsyncRequest) GetParam0() []ItemDiscountSku {
    return r.param0
}

func (r *AlibabaWdkMarketingDiscountItemRemoveAsyncRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaWdkMarketingDiscountItemRemoveAsyncRequest) GetParam1() *CommonActivityParam {
    return r.param1
}

func (r *AlibabaWdkMarketingDiscountItemRemoveAsyncRequest) SetVersion(version int64) error {
    r.version = version
    r.Set("version", version)
    return nil
}

func (r AlibabaWdkMarketingDiscountItemRemoveAsyncRequest) GetVersion() int64 {
    return r.version
}

