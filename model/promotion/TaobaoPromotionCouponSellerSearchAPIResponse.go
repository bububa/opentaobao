package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionCouponSellerSearchAPIResponse 查询绑定卖家优惠券相关信息 API返回值
// taobao.promotion.coupon.seller.search
//
// 查询绑定卖家相关优惠券信息  如isv  百川 等外部业务方
type TaobaoPromotionCouponSellerSearchAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionCouponSellerSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionCouponSellerSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionCouponSellerSearchAPIResponseModel).Reset()
}

// TaobaoPromotionCouponSellerSearchAPIResponseModel is 查询绑定卖家优惠券相关信息 成功返回结果
type TaobaoPromotionCouponSellerSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_coupon_seller_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	SellerCouponDetails []SellerCouponDetail `json:"seller_coupon_details,omitempty" xml:"seller_coupon_details>seller_coupon_detail,omitempty"`
	// 调用错误码，只有调用失败的时候才会有
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 失败详细描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 符合条件总数量，用于分页等判断
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 接口调用结果，调用成功为true，否则为false
	InvokeResult bool `json:"invoke_result,omitempty" xml:"invoke_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionCouponSellerSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SellerCouponDetails = m.SellerCouponDetails[:0]
	m.ResultCode = ""
	m.ErrorMsg = ""
	m.TotalCount = 0
	m.InvokeResult = false
}

var poolTaobaoPromotionCouponSellerSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionCouponSellerSearchAPIResponse)
	},
}

// GetTaobaoPromotionCouponSellerSearchAPIResponse 从 sync.Pool 获取 TaobaoPromotionCouponSellerSearchAPIResponse
func GetTaobaoPromotionCouponSellerSearchAPIResponse() *TaobaoPromotionCouponSellerSearchAPIResponse {
	return poolTaobaoPromotionCouponSellerSearchAPIResponse.Get().(*TaobaoPromotionCouponSellerSearchAPIResponse)
}

// ReleaseTaobaoPromotionCouponSellerSearchAPIResponse 将 TaobaoPromotionCouponSellerSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionCouponSellerSearchAPIResponse(v *TaobaoPromotionCouponSellerSearchAPIResponse) {
	v.Reset()
	poolTaobaoPromotionCouponSellerSearchAPIResponse.Put(v)
}
