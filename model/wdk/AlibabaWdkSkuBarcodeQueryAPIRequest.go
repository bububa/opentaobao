package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuBarcodeQueryAPIRequest) Reset() {
	r._skuCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuBarcodeQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.barcode.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuBarcodeQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuBarcodeQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabaWdkSkuBarcodeQueryAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabaWdkSkuBarcodeQueryAPIRequest) GetSkuCode() string {
	return r._skuCode
}

var poolAlibabaWdkSkuBarcodeQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuBarcodeQueryRequest()
	},
}

// GetAlibabaWdkSkuBarcodeQueryRequest 从 sync.Pool 获取 AlibabaWdkSkuBarcodeQueryAPIRequest
func GetAlibabaWdkSkuBarcodeQueryAPIRequest() *AlibabaWdkSkuBarcodeQueryAPIRequest {
	return poolAlibabaWdkSkuBarcodeQueryAPIRequest.Get().(*AlibabaWdkSkuBarcodeQueryAPIRequest)
}

// ReleaseAlibabaWdkSkuBarcodeQueryAPIRequest 将 AlibabaWdkSkuBarcodeQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuBarcodeQueryAPIRequest(v *AlibabaWdkSkuBarcodeQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuBarcodeQueryAPIRequest.Put(v)
}
