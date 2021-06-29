package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘专用单用户发放接口 API请求
taobao.mobile.promotion.benefit.activity.send

卖家活动中需要通过该API来发放对应的权益。手淘专用单用户发放接口。
*/
type TaobaoMobilePromotionBenefitActivitySendRequest struct {
    model.Params
    // 单用户权益发放请求
    _singleBenefitRequest   *SingleBenefitRequest
}

// 初始化TaobaoMobilePromotionBenefitActivitySendRequest对象
func NewTaobaoMobilePromotionBenefitActivitySendRequest() *TaobaoMobilePromotionBenefitActivitySendRequest{
    return &TaobaoMobilePromotionBenefitActivitySendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMobilePromotionBenefitActivitySendRequest) GetApiMethodName() string {
    return "taobao.mobile.promotion.benefit.activity.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMobilePromotionBenefitActivitySendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SingleBenefitRequest Setter
// 单用户权益发放请求
func (r *TaobaoMobilePromotionBenefitActivitySendRequest) SetSingleBenefitRequest(_singleBenefitRequest *SingleBenefitRequest) error {
    r._singleBenefitRequest = _singleBenefitRequest
    r.Set("single_benefit_request", _singleBenefitRequest)
    return nil
}

// SingleBenefitRequest Getter
func (r TaobaoMobilePromotionBenefitActivitySendRequest) GetSingleBenefitRequest() *SingleBenefitRequest {
    return r._singleBenefitRequest
}
