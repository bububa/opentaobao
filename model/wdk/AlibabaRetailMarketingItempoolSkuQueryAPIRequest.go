package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolSkuQueryAPIRequest 查询商品池活动商品【同城零售】 API请求
// alibaba.retail.marketing.itempool.sku.query
//
// 查询商品池活动商品【同城零售】
type AlibabaRetailMarketingItempoolSkuQueryAPIRequest struct {
	model.Params
	// 请求入参
	_param0 *ItemPoolActivitySkuQueryRequest
}

// NewAlibabaRetailMarketingItempoolSkuQueryRequest 初始化AlibabaRetailMarketingItempoolSkuQueryAPIRequest对象
func NewAlibabaRetailMarketingItempoolSkuQueryRequest() *AlibabaRetailMarketingItempoolSkuQueryAPIRequest {
	return &AlibabaRetailMarketingItempoolSkuQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolSkuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.sku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolSkuQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 请求入参
func (r *AlibabaRetailMarketingItempoolSkuQueryAPIRequest) SetParam0(_param0 *ItemPoolActivitySkuQueryRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaRetailMarketingItempoolSkuQueryAPIRequest) GetParam0() *ItemPoolActivitySkuQueryRequest {
	return r._param0
}
