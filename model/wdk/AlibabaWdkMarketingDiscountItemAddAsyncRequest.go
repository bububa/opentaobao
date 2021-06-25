package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
特价批量新增商品 APIRequest
alibaba.wdk.marketing.discount.item.add.async

新分组模型下新增商品
*/
type AlibabaWdkMarketingDiscountItemAddAsyncRequest struct {
    model.Params

    // sku信息
    param0   []ItemDiscountSku 

    // 系统自动生成
    param1   *CommonActivityParam 

    // alibaba.wdk.marketing.version.generate接口生成
    version   int64 

}

func NewAlibabaWdkMarketingDiscountItemAddAsyncRequest() *AlibabaWdkMarketingDiscountItemAddAsyncRequest{
    return &AlibabaWdkMarketingDiscountItemAddAsyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingDiscountItemAddAsyncRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.discount.item.add.async"
}

func (r AlibabaWdkMarketingDiscountItemAddAsyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingDiscountItemAddAsyncRequest) SetParam0(param0 []ItemDiscountSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingDiscountItemAddAsyncRequest) GetParam0() []ItemDiscountSku {
    return r.param0
}

func (r *AlibabaWdkMarketingDiscountItemAddAsyncRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaWdkMarketingDiscountItemAddAsyncRequest) GetParam1() *CommonActivityParam {
    return r.param1
}

func (r *AlibabaWdkMarketingDiscountItemAddAsyncRequest) SetVersion(version int64) error {
    r.version = version
    r.Set("version", version)
    return nil
}

func (r AlibabaWdkMarketingDiscountItemAddAsyncRequest) GetVersion() int64 {
    return r.version
}

