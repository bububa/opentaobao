package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品条码查询接口 APIRequest
alibaba.wdk.sku.barcode.query

查询商品编码，支持一品多码
*/
type AlibabaWdkSkuBarcodeQueryRequest struct {
    model.Params

    // 商品编码
    skuCode   string 

}

func NewAlibabaWdkSkuBarcodeQueryRequest() *AlibabaWdkSkuBarcodeQueryRequest{
    return &AlibabaWdkSkuBarcodeQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuBarcodeQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.barcode.query"
}

func (r AlibabaWdkSkuBarcodeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuBarcodeQueryRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

func (r AlibabaWdkSkuBarcodeQueryRequest) GetSkuCode() string {
    return r.skuCode
}

