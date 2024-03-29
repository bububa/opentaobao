package wdkitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemStoreskuQueryAPIRequest 门店商品信息查询 API请求
// alibaba.wdk.item.storesku.query
//
// 门店商品信息查询
type AlibabaWdkItemStoreskuQueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 门店ID
	_storeId string
}

// NewAlibabaWdkItemStoreskuQueryRequest 初始化AlibabaWdkItemStoreskuQueryAPIRequest对象
func NewAlibabaWdkItemStoreskuQueryRequest() *AlibabaWdkItemStoreskuQueryAPIRequest {
	return &AlibabaWdkItemStoreskuQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkItemStoreskuQueryAPIRequest) Reset() {
	r._skuCode = ""
	r._storeId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemStoreskuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.storesku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemStoreskuQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkItemStoreskuQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemStoreskuQueryAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabaWdkItemStoreskuQueryAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *AlibabaWdkItemStoreskuQueryAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaWdkItemStoreskuQueryAPIRequest) GetStoreId() string {
	return r._storeId
}

var poolAlibabaWdkItemStoreskuQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkItemStoreskuQueryRequest()
	},
}

// GetAlibabaWdkItemStoreskuQueryRequest 从 sync.Pool 获取 AlibabaWdkItemStoreskuQueryAPIRequest
func GetAlibabaWdkItemStoreskuQueryAPIRequest() *AlibabaWdkItemStoreskuQueryAPIRequest {
	return poolAlibabaWdkItemStoreskuQueryAPIRequest.Get().(*AlibabaWdkItemStoreskuQueryAPIRequest)
}

// ReleaseAlibabaWdkItemStoreskuQueryAPIRequest 将 AlibabaWdkItemStoreskuQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkItemStoreskuQueryAPIRequest(v *AlibabaWdkItemStoreskuQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkItemStoreskuQueryAPIRequest.Put(v)
}
