package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取推广计划产品列表 APIRequest
alibaba.scbp.target.ad.plan.product.list.get

定向推广-获取推广计划产品列表
*/
type AlibabaScbpTargetAdPlanProductListGetRequest struct {
    model.Params

    // TopP4pQuickProductQuery
    topP4pQuickProductQuery   *TopP4pQuickProductQuery 

}

func NewAlibabaScbpTargetAdPlanProductListGetRequest() *AlibabaScbpTargetAdPlanProductListGetRequest{
    return &AlibabaScbpTargetAdPlanProductListGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTargetAdPlanProductListGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.product.list.get"
}

func (r AlibabaScbpTargetAdPlanProductListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpTargetAdPlanProductListGetRequest) SetTopP4pQuickProductQuery(topP4pQuickProductQuery *TopP4pQuickProductQuery) error {
    r.topP4pQuickProductQuery = topP4pQuickProductQuery
    r.Set("top_p4p_quick_product_query", topP4pQuickProductQuery)
    return nil
}

func (r AlibabaScbpTargetAdPlanProductListGetRequest) GetTopP4pQuickProductQuery() *TopP4pQuickProductQuery {
    return r.topP4pQuickProductQuery
}

