package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallPromotionCouponQueryAPIResponse 查询可用优惠券列表 API返回值
// tmall.promotion.coupon.query
//
// 查询用户的可用优惠券列表，仅包含优惠券基本信息和用户nick
type TmallPromotionCouponQueryAPIResponse struct {
	model.CommonResponse
	TmallPromotionCouponQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallPromotionCouponQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallPromotionCouponQueryAPIResponseModel).Reset()
}

// TmallPromotionCouponQueryAPIResponseModel is 查询可用优惠券列表 成功返回结果
type TmallPromotionCouponQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_promotion_coupon_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallPromotionCouponQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallPromotionCouponQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallPromotionCouponQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallPromotionCouponQueryAPIResponse)
	},
}

// GetTmallPromotionCouponQueryAPIResponse 从 sync.Pool 获取 TmallPromotionCouponQueryAPIResponse
func GetTmallPromotionCouponQueryAPIResponse() *TmallPromotionCouponQueryAPIResponse {
	return poolTmallPromotionCouponQueryAPIResponse.Get().(*TmallPromotionCouponQueryAPIResponse)
}

// ReleaseTmallPromotionCouponQueryAPIResponse 将 TmallPromotionCouponQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallPromotionCouponQueryAPIResponse(v *TmallPromotionCouponQueryAPIResponse) {
	v.Reset()
	poolTmallPromotionCouponQueryAPIResponse.Put(v)
}
