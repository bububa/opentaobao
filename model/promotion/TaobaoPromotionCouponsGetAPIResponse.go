package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionCouponsGetAPIResponse 查询卖家优惠券 API返回值
// taobao.promotion.coupons.get
//
// 查询卖家已经创建的优惠券，接口返回信息：优惠券ID，面值，创建时间，有效期，使用条件，使用渠道，创建渠道，优惠券总数量
type TaobaoPromotionCouponsGetAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionCouponsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionCouponsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionCouponsGetAPIResponseModel).Reset()
}

// TaobaoPromotionCouponsGetAPIResponseModel is 查询卖家优惠券 成功返回结果
type TaobaoPromotionCouponsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_coupons_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优惠券列表
	Coupons []Coupon `json:"coupons,omitempty" xml:"coupons>coupon,omitempty"`
	// 查询的总数量
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionCouponsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Coupons = m.Coupons[:0]
	m.TotalResults = 0
}

var poolTaobaoPromotionCouponsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionCouponsGetAPIResponse)
	},
}

// GetTaobaoPromotionCouponsGetAPIResponse 从 sync.Pool 获取 TaobaoPromotionCouponsGetAPIResponse
func GetTaobaoPromotionCouponsGetAPIResponse() *TaobaoPromotionCouponsGetAPIResponse {
	return poolTaobaoPromotionCouponsGetAPIResponse.Get().(*TaobaoPromotionCouponsGetAPIResponse)
}

// ReleaseTaobaoPromotionCouponsGetAPIResponse 将 TaobaoPromotionCouponsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionCouponsGetAPIResponse(v *TaobaoPromotionCouponsGetAPIResponse) {
	v.Reset()
	poolTaobaoPromotionCouponsGetAPIResponse.Put(v)
}
