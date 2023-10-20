package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaomobilepromotioncouponsellersearchAPIResponse 查询绑定卖家优惠券相关信息(手淘专用) API返回值
// taobao.mobile.promotion.coupon.seller.search
//
// 查询绑定卖家相关优惠券信息 如isv 百川 等外部业务方
type TaobaomobilepromotioncouponsellersearchAPIResponse struct {
	model.CommonResponse
	TaobaomobilepromotioncouponsellersearchAPIResponseModel
}

// TaobaomobilepromotioncouponsellersearchAPIResponseModel is 查询绑定卖家优惠券相关信息(手淘专用) 成功返回结果
type TaobaomobilepromotioncouponsellersearchAPIResponseModel struct {
	XMLName xml.Name `xml:"mobile_promotion_coupon_seller_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优惠券查询结果
	CouponSearchResult *CouponSearchResult `json:"coupon_search_result,omitempty" xml:"coupon_search_result,omitempty"`
}
