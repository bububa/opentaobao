package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品条码查询接口 API请求
alibaba.wdk.sku.barcode.query

查询商品编码，支持一品多码
*/
type AlibabaWdkSkuBarcodeQueryRequest struct {
    model.Params
    // 商品编码
    skuCode   string
}

// 初始化AlibabaWdkSkuBarcodeQueryRequest对象
func NewAlibabaWdkSkuBarcodeQueryRequest() *AlibabaWdkSkuBarcodeQueryRequest{
    return &AlibabaWdkSkuBarcodeQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuBarcodeQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.barcode.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuBarcodeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkSkuBarcodeQueryRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkSkuBarcodeQueryRequest) GetSkuCode() string {
    return r.skuCode
}
