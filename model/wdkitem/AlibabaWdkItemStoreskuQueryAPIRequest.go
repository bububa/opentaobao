package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemstoreskuqueryAPIRequest 门店商品信息查询 API请求
// alibaba.wdk.item.storesku.query
//
// 门店商品信息查询
type AlibabawdkitemstoreskuqueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 门店ID
	_storeId string
}

// NewAlibabawdkitemstoreskuqueryRequest 初始化AlibabawdkitemstoreskuqueryAPIRequest对象
func NewAlibabawdkitemstoreskuqueryRequest() *AlibabawdkitemstoreskuqueryAPIRequest {
	return &AlibabawdkitemstoreskuqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkitemstoreskuqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.storesku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkitemstoreskuqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkitemstoreskuqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabawdkitemstoreskuqueryAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabawdkitemstoreskuqueryAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *AlibabawdkitemstoreskuqueryAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabawdkitemstoreskuqueryAPIRequest) GetStoreId() string {
	return r._storeId
}
