package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
店铺优惠券发放接口 APIResponse
taobao.promotion.coupon.send

通过接口批量发放店铺优惠券（每次只能发送100张，只能发给当前授权卖家店铺的会员），发送成功则返回为空，发送失败则返回失败的买家列表和发送成功的买家和优惠券的number。注：如果所有买家都发放失败的话，is_success也为true，建议调用者根据返回的集合判断是否送入的买家都发放成功了
*/
type TaobaoPromotionCouponSendAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"promotion_coupon_send_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // true 成功，false失败
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"