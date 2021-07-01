package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAccountBudgetGetAPIRequest
查询日消耗预算 API请求
alibaba.scbp.account.budget.get

查询日消耗预算 */
type AlibabaScbpAccountBudgetGetAPIRequest struct {
	model.Params
}

// NewAlibabaScbpAccountBudgetGetRequest 初始化AlibabaScbpAccountBudgetGetAPIRequest对象
func NewAlibabaScbpAccountBudgetGetRequest() *AlibabaScbpAccountBudgetGetAPIRequest {
	return &AlibabaScbpAccountBudgetGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAccountBudgetGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.account.budget.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAccountBudgetGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
