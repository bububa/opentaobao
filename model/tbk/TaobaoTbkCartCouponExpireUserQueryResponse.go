package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
购物车催付优惠券到期查询用户信息 APIResponse
taobao.tbk.cart.coupon.expire.user.query

购物车催付根据对应规则查询用户信息。
*/
type TaobaoTbkCartCouponExpireUserQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTbkCartCouponExpireUserQueryResponse `json:"taobao_tbk_cart_coupon_expire_user_query_response,omitempty"`
}

type TaobaoTbkCartCouponExpireUserQueryResponse struct {

    // 用户规则信息集合
    UserRuleInfoList   []UserRuleInfo `json:"user_rule_info_list,omitempty"`

}
