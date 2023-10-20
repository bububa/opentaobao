package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolActivityQueryAPIRequest 查询商品池活动【同城零售】 API请求
// alibaba.retail.marketing.itempool.activity.query
//
// 查询商品池活动【同城零售】
type AlibabaRetailMarketingItempoolActivityQueryAPIRequest struct {
	model.Params
	// 请求体
	_param0 *ItemPoolActivityQueryRequest
}

// NewAlibabaRetailMarketingItempoolActivityQueryRequest 初始化AlibabaRetailMarketingItempoolActivityQueryAPIRequest对象
func NewAlibabaRetailMarketingItempoolActivityQueryRequest() *AlibabaRetailMarketingItempoolActivityQueryAPIRequest {
	return &AlibabaRetailMarketingItempoolActivityQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolActivityQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolActivityQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingItempoolActivityQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求体
func (r *AlibabaRetailMarketingItempoolActivityQueryAPIRequest) SetParam0(_param0 *ItemPoolActivityQueryRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaRetailMarketingItempoolActivityQueryAPIRequest) GetParam0() *ItemPoolActivityQueryRequest {
	return r._param0
}
