package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpaccountbudgetgetAPIRequest 查询日消耗预算 API请求
// alibaba.scbp.account.budget.get
//
// 查询日消耗预算
type AlibabascbpaccountbudgetgetAPIRequest struct {
	model.Params
}

// NewAlibabascbpaccountbudgetgetRequest 初始化AlibabascbpaccountbudgetgetAPIRequest对象
func NewAlibabascbpaccountbudgetgetRequest() *AlibabascbpaccountbudgetgetAPIRequest {
	return &AlibabascbpaccountbudgetgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpaccountbudgetgetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.account.budget.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpaccountbudgetgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpaccountbudgetgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
