package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改关联的活动权益 API请求
taobao.promotion.benefit.activity.update

修改卖家活动中关联的对应的权益。
*/
type TaobaoPromotionBenefitActivityUpdateRequest struct {
    model.Params
    // 修改关联的权益的活动请求
    updateRequest   *UpdateBenefitActivityRequest
}

// 初始化TaobaoPromotionBenefitActivityUpdateRequest对象
func NewTaobaoPromotionBenefitActivityUpdateRequest() *TaobaoPromotionBenefitActivityUpdateRequest{
    return &TaobaoPromotionBenefitActivityUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionBenefitActivityUpdateRequest) GetApiMethodName() string {
    return "taobao.promotion.benefit.activity.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionBenefitActivityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UpdateRequest Setter
// 修改关联的权益的活动请求
func (r *TaobaoPromotionBenefitActivityUpdateRequest) SetUpdateRequest(updateRequest *UpdateBenefitActivityRequest) error {
    r.updateRequest = updateRequest
    r.Set("update_request", updateRequest)
    return nil
}

// UpdateRequest Getter
func (r TaobaoPromotionBenefitActivityUpdateRequest) GetUpdateRequest() *UpdateBenefitActivityRequest {
    return r.updateRequest
}
