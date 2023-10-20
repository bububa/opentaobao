package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemmerchantcategoryqueryAPIRequest 查询商品的商家叶子类目 API请求
// alibaba.wdk.item.merchant.category.query
//
// 查询商品的商家叶子类目
type AlibabawdkitemmerchantcategoryqueryAPIRequest struct {
	model.Params
	// 请求
	_queryRequest *WdkOpenSkuMerchantCatServiceQueryRequest
}

// NewAlibabawdkitemmerchantcategoryqueryRequest 初始化AlibabawdkitemmerchantcategoryqueryAPIRequest对象
func NewAlibabawdkitemmerchantcategoryqueryRequest() *AlibabawdkitemmerchantcategoryqueryAPIRequest {
	return &AlibabawdkitemmerchantcategoryqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkitemmerchantcategoryqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.merchant.category.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkitemmerchantcategoryqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkitemmerchantcategoryqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryRequest is QueryRequest Setter
// 请求
func (r *AlibabawdkitemmerchantcategoryqueryAPIRequest) SetQueryRequest(_queryRequest *WdkOpenSkuMerchantCatServiceQueryRequest) error {
	r._queryRequest = _queryRequest
	r.Set("query_request", _queryRequest)
	return nil
}

// GetQueryRequest QueryRequest Getter
func (r AlibabawdkitemmerchantcategoryqueryAPIRequest) GetQueryRequest() *WdkOpenSkuMerchantCatServiceQueryRequest {
	return r._queryRequest
}
