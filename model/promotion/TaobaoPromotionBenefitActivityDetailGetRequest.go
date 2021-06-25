package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
活动关联的权益详情获取 APIRequest
taobao.promotion.benefit.activity.detail.get

活动关联的权益详情获取
*/
type TaobaoPromotionBenefitActivityDetailGetRequest struct {
    model.Params

    // 查询活动关联权益详情的请求
    queryRequest   *ActivityRelationDetailRequest 

}

func NewTaobaoPromotionBenefitActivityDetailGetRequest() *TaobaoPromotionBenefitActivityDetailGetRequest{
    return &TaobaoPromotionBenefitActivityDetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionBenefitActivityDetailGetRequest) GetApiMethodName() string {
    return "taobao.promotion.benefit.activity.detail.get"
}

func (r TaobaoPromotionBenefitActivityDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionBenefitActivityDetailGetRequest) SetQueryRequest(queryRequest *ActivityRelationDetailRequest) error {
    r.queryRequest = queryRequest
    r.Set("query_request", queryRequest)
    return nil
}

func (r TaobaoPromotionBenefitActivityDetailGetRequest) GetQueryRequest() *ActivityRelationDetailRequest {
    return r.queryRequest
}

