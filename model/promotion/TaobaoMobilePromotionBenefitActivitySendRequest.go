package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘专用单用户发放接口 APIRequest
taobao.mobile.promotion.benefit.activity.send

卖家活动中需要通过该API来发放对应的权益。手淘专用单用户发放接口。
*/
type TaobaoMobilePromotionBenefitActivitySendRequest struct {
    model.Params

    // 单用户权益发放请求
    singleBenefitRequest   *SingleBenefitRequest 

}

func NewTaobaoMobilePromotionBenefitActivitySendRequest() *TaobaoMobilePromotionBenefitActivitySendRequest{
    return &TaobaoMobilePromotionBenefitActivitySendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMobilePromotionBenefitActivitySendRequest) GetApiMethodName() string {
    return "taobao.mobile.promotion.benefit.activity.send"
}

func (r TaobaoMobilePromotionBenefitActivitySendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMobilePromotionBenefitActivitySendRequest) SetSingleBenefitRequest(singleBenefitRequest *SingleBenefitRequest) error {
    r.singleBenefitRequest = singleBenefitRequest
    r.Set("single_benefit_request", singleBenefitRequest)
    return nil
}

func (r TaobaoMobilePromotionBenefitActivitySendRequest) GetSingleBenefitRequest() *SingleBenefitRequest {
    return r.singleBenefitRequest
}

