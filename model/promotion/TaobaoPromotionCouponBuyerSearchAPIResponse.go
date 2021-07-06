package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionCouponBuyerSearchAPIResponse 查询买家在相关app领取的优惠券信息 API返回值
// taobao.promotion.coupon.buyer.search
//
// 查询买家在相关app领取的优惠券信息
type TaobaoPromotionCouponBuyerSearchAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionCouponBuyerSearchAPIResponseModel
}

// TaobaoPromotionCouponBuyerSearchAPIResponseModel is 查询买家在相关app领取的优惠券信息 成功返回结果
type TaobaoPromotionCouponBuyerSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_coupon_buyer_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	BuyerCouponInfos []BuyerCouponInfo `json:"buyer_coupon_infos,omitempty" xml:"buyer_coupon_infos>buyer_coupon_info,omitempty"`
	// 结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 符合条件的总数，用于分页判断
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 调用是否成功
	InvokeResult bool `json:"invoke_result,omitempty" xml:"invoke_result,omitempty"`
}
