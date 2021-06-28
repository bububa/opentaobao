package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询绑定卖家优惠券相关信息 APIResponse
taobao.promotion.coupon.seller.search

查询绑定卖家相关优惠券信息  如isv  百川 等外部业务方
*/
type TaobaoPromotionCouponSellerSearchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"promotion_coupon_seller_search_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用错误码，只有调用失败的时候才会有
    
    ResultCode   string `json:"result_code,omitempty" xml:"