package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingitemdiscountactivitysaveAPIRequest 【同城零售】单品活动保存 API请求
// alibaba.retail.marketing.itemdiscount.activity.save
//
// 【同城零售】单品活动保存
type AlibabaretailmarketingitemdiscountactivitysaveAPIRequest struct {
	model.Params
	// 保存活动参数
	_param *ItemDiscountActivityOperateRequest
}

// NewAlibabaretailmarketingitemdiscountactivitysaveRequest 初始化AlibabaretailmarketingitemdiscountactivitysaveAPIRequest对象
func NewAlibabaretailmarketingitemdiscountactivitysaveRequest() *AlibabaretailmarketingitemdiscountactivitysaveAPIRequest {
	return &AlibabaretailmarketingitemdiscountactivitysaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingitemdiscountactivitysaveAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.activity.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingitemdiscountactivitysaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingitemdiscountactivitysaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 保存活动参数
func (r *AlibabaretailmarketingitemdiscountactivitysaveAPIRequest) SetParam(_param *ItemDiscountActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaretailmarketingitemdiscountactivitysaveAPIRequest) GetParam() *ItemDiscountActivityOperateRequest {
	return r._param
}
