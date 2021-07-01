package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取推广计划产品列表 API请求
alibaba.scbp.target.ad.plan.product.list.get

定向推广-获取推广计划产品列表
*/
type AlibabaScbpTargetAdPlanProductListGetAPIRequest struct {
    model.Params
    // TopP4pQuickProductQuery
    _topP4pQuickProductQuery   *TopP4pQuickProductQuery
}

// 初始化AlibabaScbpTargetAdPlanProductListGetAPIRequest对象
func NewAlibabaScbpTargetAdPlanProductListGetRequest() *AlibabaScbpTargetAdPlanProductListGetAPIRequest{
    return &AlibabaScbpTargetAdPlanProductListGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanProductListGetAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.product.list.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanProductListGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopP4pQuickProductQuery Setter
// TopP4pQuickProductQuery
func (r *AlibabaScbpTargetAdPlanProductListGetAPIRequest) SetTopP4pQuickProductQuery(_topP4pQuickProductQuery *TopP4pQuickProductQuery) error {
    r._topP4pQuickProductQuery = _topP4pQuickProductQuery
    r.Set("top_p4p_quick_product_query", _topP4pQuickProductQuery)
    return nil
}

// TopP4pQuickProductQuery Getter
func (r AlibabaScbpTargetAdPlanProductListGetAPIRequest) GetTopP4pQuickProductQuery() *TopP4pQuickProductQuery {
    return r._topP4pQuickProductQuery
}
