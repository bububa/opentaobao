package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemMerchantstoreskuCreateAPIRequest
门店商品信息新建 API请求
alibaba.wdk.item.merchantstoresku.create

门店商品信息新建 */
type AlibabaWdkItemMerchantstoreskuCreateAPIRequest struct {
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

// NewAlibabaWdkItemMerchantstoreskuCreateRequest 初始化AlibabaWdkItemMerchantstoreskuCreateAPIRequest对象
func NewAlibabaWdkItemMerchantstoreskuCreateRequest() *AlibabaWdkItemMerchantstoreskuCreateAPIRequest {
	return &AlibabaWdkItemMerchantstoreskuCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMerchantstoreskuCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.merchantstoresku.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMerchantstoreskuCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreId Setter
// 门店编码
func (r *AlibabaWdkItemMerchantstoreskuCreateAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaWdkItemMerchantstoreskuCreateAPIRequest) GetStoreId() string {
	return r._storeId
}

// Set is SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemMerchantstoreskuCreateAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// Get SkuCode Getter
func (r AlibabaWdkItemMerchantstoreskuCreateAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// Set is Params Setter
// 新建参数，是个json字符串
func (r *AlibabaWdkItemMerchantstoreskuCreateAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// Get Params Getter
func (r AlibabaWdkItemMerchantstoreskuCreateAPIRequest) GetParams() string {
	return r._params
}

// Set is OrgCode Setter
// 机构
func (r *AlibabaWdkItemMerchantstoreskuCreateAPIRequest) SetOrgCode(_orgCode string) error {
	r._orgCode = _orgCode
	r.Set("org_code", _orgCode)
	return nil
}

// Get OrgCode Getter
func (r AlibabaWdkItemMerchantstoreskuCreateAPIRequest) GetOrgCode() string {
	return r._orgCode
}
