package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingitempoolskuqueryAPIRequest 查询商品池活动商品【同城零售】 API请求
// alibaba.retail.marketing.itempool.sku.query
//
// 查询商品池活动商品【同城零售】
type AlibabaretailmarketingitempoolskuqueryAPIRequest struct {
	model.Params
	// 请求入参
	_param0 *ItemPoolActivitySkuQueryRequest
}

// NewAlibabaretailmarketingitempoolskuqueryRequest 初始化AlibabaretailmarketingitempoolskuqueryAPIRequest对象
func NewAlibabaretailmarketingitempoolskuqueryRequest() *AlibabaretailmarketingitempoolskuqueryAPIRequest {
	return &AlibabaretailmarketingitempoolskuqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingitempoolskuqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.sku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingitempoolskuqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingitempoolskuqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求入参
func (r *AlibabaretailmarketingitempoolskuqueryAPIRequest) SetParam0(_param0 *ItemPoolActivitySkuQueryRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaretailmarketingitempoolskuqueryAPIRequest) GetParam0() *ItemPoolActivitySkuQueryRequest {
	return r._param0
}
