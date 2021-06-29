package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
活动关联的权益详情获取 API请求
taobao.promotion.benefit.activity.detail.get

活动关联的权益详情获取
*/
type TaobaoPromotionBenefitActivityDetailGetRequest struct {
    model.Params
    // 查询活动关联权益详情的请求
    _queryRequest   *ActivityRelationDetailRequest
}

// 初始化TaobaoPromotionBenefitActivityDetailGetRequest对象
func NewTaobaoPromotionBenefitActivityDetailGetRequest() *TaobaoPromotionBenefitActivityDetailGetRequest{
    return &TaobaoPromotionBenefitActivityDetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionBenefitActivityDetailGetRequest) GetApiMethodName() string {
    return "taobao.promotion.benefit.activity.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionBenefitActivityDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryRequest Setter
// 查询活动关联权益详情的请求
func (r *TaobaoPromotionBenefitActivityDetailGetRequest) SetQueryRequest(_queryRequest *ActivityRelationDetailRequest) error {
    r._queryRequest = _queryRequest
    r.Set("query_request", _queryRequest)
    return nil
}

// QueryRequest Getter
func (r TaobaoPromotionBenefitActivityDetailGetRequest) GetQueryRequest() *ActivityRelationDetailRequest {
    return r._queryRequest
}
