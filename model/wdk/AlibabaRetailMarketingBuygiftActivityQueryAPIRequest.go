package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingbuygiftactivityqueryAPIRequest 查询单品买赠活动【同城零售】 API请求
// alibaba.retail.marketing.buygift.activity.query
//
// 查询单品买赠活动【同城零售】
type AlibabaretailmarketingbuygiftactivityqueryAPIRequest struct {
	model.Params
	// 买赠活动查询入参
	_param0 *BuyGiftActivityQueryRequest
}

// NewAlibabaretailmarketingbuygiftactivityqueryRequest 初始化AlibabaretailmarketingbuygiftactivityqueryAPIRequest对象
func NewAlibabaretailmarketingbuygiftactivityqueryRequest() *AlibabaretailmarketingbuygiftactivityqueryAPIRequest {
	return &AlibabaretailmarketingbuygiftactivityqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingbuygiftactivityqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingbuygiftactivityqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingbuygiftactivityqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 买赠活动查询入参
func (r *AlibabaretailmarketingbuygiftactivityqueryAPIRequest) SetParam0(_param0 *BuyGiftActivityQueryRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaretailmarketingbuygiftactivityqueryAPIRequest) GetParam0() *BuyGiftActivityQueryRequest {
	return r._param0
}
