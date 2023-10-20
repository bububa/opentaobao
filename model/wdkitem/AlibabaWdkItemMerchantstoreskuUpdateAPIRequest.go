package wdkitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemMerchantstoreskuUpdateAPIRequest 门店商品信息修改 API请求
// alibaba.wdk.item.merchantstoresku.update
//
// 门店商品信息修改
type AlibabaWdkItemMerchantstoreskuUpdateAPIRequest struct {
	model.Params
	// 门店ID
	_storeId string
	// 商品编码
	_skuCode string
	// 修改参数，是个json字符串
	_params string
}

// NewAlibabaWdkItemMerchantstoreskuUpdateRequest 初始化AlibabaWdkItemMerchantstoreskuUpdateAPIRequest对象
func NewAlibabaWdkItemMerchantstoreskuUpdateRequest() *AlibabaWdkItemMerchantstoreskuUpdateAPIRequest {
	return &AlibabaWdkItemMerchantstoreskuUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) Reset() {
	r._storeId = ""
	r._skuCode = ""
	r._params = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.merchantstoresku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetParams is Params Setter
// 修改参数，是个json字符串
func (r *AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) GetParams() string {
	return r._params
}

var poolAlibabaWdkItemMerchantstoreskuUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkItemMerchantstoreskuUpdateRequest()
	},
}

// GetAlibabaWdkItemMerchantstoreskuUpdateRequest 从 sync.Pool 获取 AlibabaWdkItemMerchantstoreskuUpdateAPIRequest
func GetAlibabaWdkItemMerchantstoreskuUpdateAPIRequest() *AlibabaWdkItemMerchantstoreskuUpdateAPIRequest {
	return poolAlibabaWdkItemMerchantstoreskuUpdateAPIRequest.Get().(*AlibabaWdkItemMerchantstoreskuUpdateAPIRequest)
}

// ReleaseAlibabaWdkItemMerchantstoreskuUpdateAPIRequest 将 AlibabaWdkItemMerchantstoreskuUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkItemMerchantstoreskuUpdateAPIRequest(v *AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) {
	v.Reset()
	poolAlibabaWdkItemMerchantstoreskuUpdateAPIRequest.Put(v)
}
