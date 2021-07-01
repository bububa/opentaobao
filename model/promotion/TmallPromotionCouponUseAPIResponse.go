package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
券核销接口 API返回值 
tmall.promotion.coupon.use

核销用户的一张优惠券，返回核销结果
*/
type TmallPromotionCouponUseAPIResponse struct {
    model.CommonResponse
    TmallPromotionCouponUseAPIResponseModel
}

// 券核销接口 成功返回结果
type TmallPromotionCouponUseAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_promotion_coupon_use_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // data
    Data   *UseResultDo `json:"data,omitempty" xml:"data,omitempty"`
    // resultCode
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
