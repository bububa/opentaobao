package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMobilePromotionCouponSellerSearchAPIResponse 查询绑定卖家优惠券相关信息(手淘专用) API返回值
// taobao.mobile.promotion.coupon.seller.search
//
// 查询绑定卖家相关优惠券信息 如isv 百川 等外部业务方
type TaobaoMobilePromotionCouponSellerSearchAPIResponse struct {
	model.CommonResponse
	TaobaoMobilePromotionCouponSellerSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMobilePromotionCouponSellerSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMobilePromotionCouponSellerSearchAPIResponseModel).Reset()
}

// TaobaoMobilePromotionCouponSellerSearchAPIResponseModel is 查询绑定卖家优惠券相关信息(手淘专用) 成功返回结果
type TaobaoMobilePromotionCouponSellerSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"mobile_promotion_coupon_seller_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优惠券查询结果
	CouponSearchResult *CouponSearchResult `json:"coupon_search_result,omitempty" xml:"coupon_search_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMobilePromotionCouponSellerSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CouponSearchResult = nil
}

var poolTaobaoMobilePromotionCouponSellerSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMobilePromotionCouponSellerSearchAPIResponse)
	},
}

// GetTaobaoMobilePromotionCouponSellerSearchAPIResponse 从 sync.Pool 获取 TaobaoMobilePromotionCouponSellerSearchAPIResponse
func GetTaobaoMobilePromotionCouponSellerSearchAPIResponse() *TaobaoMobilePromotionCouponSellerSearchAPIResponse {
	return poolTaobaoMobilePromotionCouponSellerSearchAPIResponse.Get().(*TaobaoMobilePromotionCouponSellerSearchAPIResponse)
}

// ReleaseTaobaoMobilePromotionCouponSellerSearchAPIResponse 将 TaobaoMobilePromotionCouponSellerSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMobilePromotionCouponSellerSearchAPIResponse(v *TaobaoMobilePromotionCouponSellerSearchAPIResponse) {
	v.Reset()
	poolTaobaoMobilePromotionCouponSellerSearchAPIResponse.Put(v)
}
