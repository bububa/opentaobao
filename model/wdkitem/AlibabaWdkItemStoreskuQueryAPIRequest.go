package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemStoreskuQueryAPIRequest
门店商品信息查询 API请求
alibaba.wdk.item.storesku.query

门店商品信息查询 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemStoreskuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.storesku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemStoreskuQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemStoreskuQueryAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// Get SkuCode Getter
func (r AlibabaWdkItemStoreskuQueryAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// Set is StoreId Setter
// 门店ID
func (r *AlibabaWdkItemStoreskuQueryAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaWdkItemStoreskuQueryAPIRequest) GetStoreId() string {
	return r._storeId
}
