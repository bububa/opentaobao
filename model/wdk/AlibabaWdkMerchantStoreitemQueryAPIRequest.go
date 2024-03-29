package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantStoreitemQueryAPIRequest 门店商品信心查询 API请求
// alibaba.wdk.merchant.storeitem.query
//
// 门店商品信心查询
type AlibabaWdkMerchantStoreitemQueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 商家编码
	_merchantCode string
	// 门店编码
	_storeId string
}

// NewAlibabaWdkMerchantStoreitemQueryRequest 初始化AlibabaWdkMerchantStoreitemQueryAPIRequest对象
func NewAlibabaWdkMerchantStoreitemQueryRequest() *AlibabaWdkMerchantStoreitemQueryAPIRequest {
	return &AlibabaWdkMerchantStoreitemQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMerchantStoreitemQueryAPIRequest) Reset() {
	r._skuCode = ""
	r._merchantCode = ""
	r._storeId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantStoreitemQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchant.storeitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantStoreitemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMerchantStoreitemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabaWdkMerchantStoreitemQueryAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabaWdkMerchantStoreitemQueryAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetMerchantCode is MerchantCode Setter
// 商家编码
func (r *AlibabaWdkMerchantStoreitemQueryAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// GetMerchantCode MerchantCode Getter
func (r AlibabaWdkMerchantStoreitemQueryAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

// SetStoreId is StoreId Setter
// 门店编码
func (r *AlibabaWdkMerchantStoreitemQueryAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaWdkMerchantStoreitemQueryAPIRequest) GetStoreId() string {
	return r._storeId
}

var poolAlibabaWdkMerchantStoreitemQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMerchantStoreitemQueryRequest()
	},
}

// GetAlibabaWdkMerchantStoreitemQueryRequest 从 sync.Pool 获取 AlibabaWdkMerchantStoreitemQueryAPIRequest
func GetAlibabaWdkMerchantStoreitemQueryAPIRequest() *AlibabaWdkMerchantStoreitemQueryAPIRequest {
	return poolAlibabaWdkMerchantStoreitemQueryAPIRequest.Get().(*AlibabaWdkMerchantStoreitemQueryAPIRequest)
}

// ReleaseAlibabaWdkMerchantStoreitemQueryAPIRequest 将 AlibabaWdkMerchantStoreitemQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMerchantStoreitemQueryAPIRequest(v *AlibabaWdkMerchantStoreitemQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkMerchantStoreitemQueryAPIRequest.Put(v)
}
