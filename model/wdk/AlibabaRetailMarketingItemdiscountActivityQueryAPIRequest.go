package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingitemdiscountactivityqueryAPIRequest 查询单品特价活动【同城零售】 API请求
// alibaba.retail.marketing.itemdiscount.activity.query
//
// 查询单品特价活动【同城零售】
type AlibabaretailmarketingitemdiscountactivityqueryAPIRequest struct {
	model.Params
	// 请求体
	_param0 *ItemDiscountActivityQueryRequest
}

// NewAlibabaretailmarketingitemdiscountactivityqueryRequest 初始化AlibabaretailmarketingitemdiscountactivityqueryAPIRequest对象
func NewAlibabaretailmarketingitemdiscountactivityqueryRequest() *AlibabaretailmarketingitemdiscountactivityqueryAPIRequest {
	return &AlibabaretailmarketingitemdiscountactivityqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingitemdiscountactivityqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingitemdiscountactivityqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingitemdiscountactivityqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求体
func (r *AlibabaretailmarketingitemdiscountactivityqueryAPIRequest) SetParam0(_param0 *ItemDiscountActivityQueryRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaretailmarketingitemdiscountactivityqueryAPIRequest) GetParam0() *ItemDiscountActivityQueryRequest {
	return r._param0
}
