package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemmerchantstoreskucreateAPIRequest 门店商品信息新建 API请求
// alibaba.wdk.item.merchantstoresku.create
//
// 门店商品信息新建
type AlibabawdkitemmerchantstoreskucreateAPIRequest struct {
	model.Params
	// 门店编码
	_storeId string
	// 商品编码
	_skuCode string
	// 新建参数，是个json字符串
	_params string
	// 机构
	_orgCode string
}

// NewAlibabawdkitemmerchantstoreskucreateRequest 初始化AlibabawdkitemmerchantstoreskucreateAPIRequest对象
func NewAlibabawdkitemmerchantstoreskucreateRequest() *AlibabawdkitemmerchantstoreskucreateAPIRequest {
	return &AlibabawdkitemmerchantstoreskucreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkitemmerchantstoreskucreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.merchantstoresku.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkitemmerchantstoreskucreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkitemmerchantstoreskucreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店编码
func (r *AlibabawdkitemmerchantstoreskucreateAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabawdkitemmerchantstoreskucreateAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabawdkitemmerchantstoreskucreateAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabawdkitemmerchantstoreskucreateAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetParams is Params Setter
// 新建参数，是个json字符串
func (r *AlibabawdkitemmerchantstoreskucreateAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r AlibabawdkitemmerchantstoreskucreateAPIRequest) GetParams() string {
	return r._params
}

// SetOrgCode is OrgCode Setter
// 机构
func (r *AlibabawdkitemmerchantstoreskucreateAPIRequest) SetOrgCode(_orgCode string) error {
	r._orgCode = _orgCode
	r.Set("org_code", _orgCode)
	return nil
}

// GetOrgCode OrgCode Getter
func (r AlibabawdkitemmerchantstoreskucreateAPIRequest) GetOrgCode() string {
	return r._orgCode
}
