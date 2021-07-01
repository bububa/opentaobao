package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTargetAdPlanCountryIdGetAPIRequest
定向推广-国家标签ID获取 API请求
alibaba.scbp.target.ad.plan.country.id.get

定向推广-国家标签ID获取 */
type AlibabaScbpTargetAdPlanCountryIdGetAPIRequest struct {
	model.Params
}

// NewAlibabaScbpTargetAdPlanCountryIdGetRequest 初始化AlibabaScbpTargetAdPlanCountryIdGetAPIRequest对象
func NewAlibabaScbpTargetAdPlanCountryIdGetRequest() *AlibabaScbpTargetAdPlanCountryIdGetAPIRequest {
	return &AlibabaScbpTargetAdPlanCountryIdGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanCountryIdGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.country.id.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanCountryIdGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
