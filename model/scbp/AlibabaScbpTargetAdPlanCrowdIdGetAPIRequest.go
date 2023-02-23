package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanCrowdIdGetAPIRequest 定向推广-人群标签ID获取(店铺老客、优选人群) API请求
// alibaba.scbp.target.ad.plan.crowd.id.get
//
// 定向推广-人群标签ID获取(店铺老客、优选人群)
type AlibabaScbpTargetAdPlanCrowdIdGetAPIRequest struct {
	model.Params
}

// NewAlibabaScbpTargetAdPlanCrowdIdGetRequest 初始化AlibabaScbpTargetAdPlanCrowdIdGetAPIRequest对象
func NewAlibabaScbpTargetAdPlanCrowdIdGetRequest() *AlibabaScbpTargetAdPlanCrowdIdGetAPIRequest {
	return &AlibabaScbpTargetAdPlanCrowdIdGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanCrowdIdGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.crowd.id.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanCrowdIdGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpTargetAdPlanCrowdIdGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
