package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanCountryIdGetAPIRequest 定向推广-国家标签ID获取 API请求
// alibaba.scbp.target.ad.plan.country.id.get
//
// 定向推广-国家标签ID获取
type AlibabaScbpTargetAdPlanCountryIdGetAPIRequest struct {
	model.Params
}

// NewAlibabaScbpTargetAdPlanCountryIdGetRequest 初始化AlibabaScbpTargetAdPlanCountryIdGetAPIRequest对象
func NewAlibabaScbpTargetAdPlanCountryIdGetRequest() *AlibabaScbpTargetAdPlanCountryIdGetAPIRequest {
	return &AlibabaScbpTargetAdPlanCountryIdGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpTargetAdPlanCountryIdGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanCountryIdGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.country.id.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanCountryIdGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpTargetAdPlanCountryIdGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaScbpTargetAdPlanCountryIdGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpTargetAdPlanCountryIdGetRequest()
	},
}

// GetAlibabaScbpTargetAdPlanCountryIdGetRequest 从 sync.Pool 获取 AlibabaScbpTargetAdPlanCountryIdGetAPIRequest
func GetAlibabaScbpTargetAdPlanCountryIdGetAPIRequest() *AlibabaScbpTargetAdPlanCountryIdGetAPIRequest {
	return poolAlibabaScbpTargetAdPlanCountryIdGetAPIRequest.Get().(*AlibabaScbpTargetAdPlanCountryIdGetAPIRequest)
}

// ReleaseAlibabaScbpTargetAdPlanCountryIdGetAPIRequest 将 AlibabaScbpTargetAdPlanCountryIdGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpTargetAdPlanCountryIdGetAPIRequest(v *AlibabaScbpTargetAdPlanCountryIdGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpTargetAdPlanCountryIdGetAPIRequest.Put(v)
}
