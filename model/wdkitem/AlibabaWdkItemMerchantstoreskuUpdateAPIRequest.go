package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemmerchantstoreskuupdateAPIRequest 门店商品信息修改 API请求
// alibaba.wdk.item.merchantstoresku.update
//
// 门店商品信息修改
type AlibabawdkitemmerchantstoreskuupdateAPIRequest struct {
	model.Params
	// 门店ID
	_storeId string
	// 商品编码
	_skuCode string
	// 修改参数，是个json字符串
	_params string
}

// NewAlibabawdkitemmerchantstoreskuupdateRequest 初始化AlibabawdkitemmerchantstoreskuupdateAPIRequest对象
func NewAlibabawdkitemmerchantstoreskuupdateRequest() *AlibabawdkitemmerchantstoreskuupdateAPIRequest {
	return &AlibabawdkitemmerchantstoreskuupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkitemmerchantstoreskuupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.merchantstoresku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkitemmerchantstoreskuupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkitemmerchantstoreskuupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *AlibabawdkitemmerchantstoreskuupdateAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabawdkitemmerchantstoreskuupdateAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabawdkitemmerchantstoreskuupdateAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabawdkitemmerchantstoreskuupdateAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetParams is Params Setter
// 修改参数，是个json字符串
func (r *AlibabawdkitemmerchantstoreskuupdateAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r AlibabawdkitemmerchantstoreskuupdateAPIRequest) GetParams() string {
	return r._params
}
