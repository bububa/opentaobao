package tmalltrend

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallBrandItemCouponProtectAPIResponse 全域新品店铺优惠券免除 API返回值
// tmall.brand.item.coupon.protect
//
// 全域新品店铺优惠券免除申请打标接口
type TmallBrandItemCouponProtectAPIResponse struct {
	model.CommonResponse
	TmallBrandItemCouponProtectAPIResponseModel
}

// TmallBrandItemCouponProtectAPIResponseModel is 全域新品店铺优惠券免除 成功返回结果
type TmallBrandItemCouponProtectAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_brand_item_coupon_protect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 店铺优惠券保护期设置是否成功
	RespSuccess bool `json:"resp_success,omitempty" xml:"resp_success,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	RespErrorCode string `json:"resp_error_code,omitempty" xml:"resp_error_code,omitempty"`
}
