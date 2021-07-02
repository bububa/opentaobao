package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuBarcodeQueryAPIRequest 商品条码查询接口 API请求
// alibaba.wdk.sku.barcode.query
//
// 查询商品编码，支持一品多码
type AlibabaWdkSkuBarcodeQueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
}

// NewAlibabaWdkSkuBarcodeQueryRequest 初始化AlibabaWdkSkuBarcodeQueryAPIRequest对象
func NewAlibabaWdkSkuBarcodeQueryRequest() *AlibabaWdkSkuBarcodeQueryAPIRequest {
	return &AlibabaWdkSkuBarcodeQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuBarcodeQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.barcode.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuBarcodeQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SkuCode Setter
// 商品编码
func (r *AlibabaWdkSkuBarcodeQueryAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// Get SkuCode Getter
func (r AlibabaWdkSkuBarcodeQueryAPIRequest) GetSkuCode() string {
	return r._skuCode
}
