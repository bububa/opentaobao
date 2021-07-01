package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkCartCouponExpireUserQueryAPIResponse
购物车催付优惠券到期查询用户信息 API返回值
taobao.tbk.cart.coupon.expire.user.query

购物车催付根据对应规则查询用户信息。 */
type TaobaoTbkCartCouponExpireUserQueryAPIResponse struct {
	model.CommonResponse
	TaobaoTbkCartCouponExpireUserQueryAPIResponseModel
}

// TaobaoTbkCartCouponExpireUserQueryAPIResponseModel is 购物车催付优惠券到期查询用户信息 成功返回结果
type TaobaoTbkCartCouponExpireUserQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_cart_coupon_expire_user_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户规则信息集合
	UserRuleInfoList []UserRuleInfo `json:"user_rule_info_list,omitempty" xml:"user_rule_info_list>user_rule_info,omitempty"`
}
