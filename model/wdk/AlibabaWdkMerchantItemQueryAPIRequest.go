package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantItemQueryAPIRequest 商家商品查询 API请求
// alibaba.wdk.merchant.item.query
//
// 商家商品查询
type AlibabaWdkMerchantItemQueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 商家编码
	_merchantCode string
}

// NewAlibabaWdkMerchantItemQueryRequest 初始化AlibabaWdkMerchantItemQueryAPIRequest对象
func NewAlibabaWdkMerchantItemQueryRequest() *AlibabaWdkMerchantItemQueryAPIRequest {
	return &AlibabaWdkMerchantItemQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMerchantItemQueryAPIRequest) Reset() {
	r._skuCode = ""
	r._merchantCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantItemQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchant.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantItemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMerchantItemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabaWdkMerchantItemQueryAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabaWdkMerchantItemQueryAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetMerchantCode is MerchantCode Setter
// 商家编码
func (r *AlibabaWdkMerchantItemQueryAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// GetMerchantCode MerchantCode Getter
func (r AlibabaWdkMerchantItemQueryAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

var poolAlibabaWdkMerchantItemQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMerchantItemQueryRequest()
	},
}

// GetAlibabaWdkMerchantItemQueryRequest 从 sync.Pool 获取 AlibabaWdkMerchantItemQueryAPIRequest
func GetAlibabaWdkMerchantItemQueryAPIRequest() *AlibabaWdkMerchantItemQueryAPIRequest {
	return poolAlibabaWdkMerchantItemQueryAPIRequest.Get().(*AlibabaWdkMerchantItemQueryAPIRequest)
}

// ReleaseAlibabaWdkMerchantItemQueryAPIRequest 将 AlibabaWdkMerchantItemQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMerchantItemQueryAPIRequest(v *AlibabaWdkMerchantItemQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkMerchantItemQueryAPIRequest.Put(v)
}
