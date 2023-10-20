package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingitempoolactivityqueryAPIRequest 查询商品池活动【同城零售】 API请求
// alibaba.retail.marketing.itempool.activity.query
//
// 查询商品池活动【同城零售】
type AlibabaretailmarketingitempoolactivityqueryAPIRequest struct {
	model.Params
	// 请求体
	_param0 *ItemPoolActivityQueryRequest
}

// NewAlibabaretailmarketingitempoolactivityqueryRequest 初始化AlibabaretailmarketingitempoolactivityqueryAPIRequest对象
func NewAlibabaretailmarketingitempoolactivityqueryRequest() *AlibabaretailmarketingitempoolactivityqueryAPIRequest {
	return &AlibabaretailmarketingitempoolactivityqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingitempoolactivityqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingitempoolactivityqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingitempoolactivityqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求体
func (r *AlibabaretailmarketingitempoolactivityqueryAPIRequest) SetParam0(_param0 *ItemPoolActivityQueryRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaretailmarketingitempoolactivityqueryAPIRequest) GetParam0() *ItemPoolActivityQueryRequest {
	return r._param0
}
