package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanProductListGetAPIRequest 定向推广-获取推广计划产品列表 API请求
// alibaba.scbp.target.ad.plan.product.list.get
//
// 定向推广-获取推广计划产品列表
type AlibabaScbpTargetAdPlanProductListGetAPIRequest struct {
	model.Params
	// TopP4pQuickProductQuery
	_topP4pQuickProductQuery *TopP4pQuickProductQuery
}

// NewAlibabaScbpTargetAdPlanProductListGetRequest 初始化AlibabaScbpTargetAdPlanProductListGetAPIRequest对象
func NewAlibabaScbpTargetAdPlanProductListGetRequest() *AlibabaScbpTargetAdPlanProductListGetAPIRequest {
	return &AlibabaScbpTargetAdPlanProductListGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpTargetAdPlanProductListGetAPIRequest) Reset() {
	r._topP4pQuickProductQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanProductListGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.product.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanProductListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpTargetAdPlanProductListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopP4pQuickProductQuery is TopP4pQuickProductQuery Setter
// TopP4pQuickProductQuery
func (r *AlibabaScbpTargetAdPlanProductListGetAPIRequest) SetTopP4pQuickProductQuery(_topP4pQuickProductQuery *TopP4pQuickProductQuery) error {
	r._topP4pQuickProductQuery = _topP4pQuickProductQuery
	r.Set("top_p4p_quick_product_query", _topP4pQuickProductQuery)
	return nil
}

// GetTopP4pQuickProductQuery TopP4pQuickProductQuery Getter
func (r AlibabaScbpTargetAdPlanProductListGetAPIRequest) GetTopP4pQuickProductQuery() *TopP4pQuickProductQuery {
	return r._topP4pQuickProductQuery
}

var poolAlibabaScbpTargetAdPlanProductListGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpTargetAdPlanProductListGetRequest()
	},
}

// GetAlibabaScbpTargetAdPlanProductListGetRequest 从 sync.Pool 获取 AlibabaScbpTargetAdPlanProductListGetAPIRequest
func GetAlibabaScbpTargetAdPlanProductListGetAPIRequest() *AlibabaScbpTargetAdPlanProductListGetAPIRequest {
	return poolAlibabaScbpTargetAdPlanProductListGetAPIRequest.Get().(*AlibabaScbpTargetAdPlanProductListGetAPIRequest)
}

// ReleaseAlibabaScbpTargetAdPlanProductListGetAPIRequest 将 AlibabaScbpTargetAdPlanProductListGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpTargetAdPlanProductListGetAPIRequest(v *AlibabaScbpTargetAdPlanProductListGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpTargetAdPlanProductListGetAPIRequest.Put(v)
}
