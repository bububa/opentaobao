package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomecouponsubmitAPIResponse 提交专车优惠券活动 API返回值
// alibaba.alihouse.newhome.coupon.submit
//
// 提交专车优惠券活动
type AlibabaalihousenewhomecouponsubmitAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomecouponsubmitAPIResponseModel
}

// AlibabaalihousenewhomecouponsubmitAPIResponseModel is 提交专车优惠券活动 成功返回结果
type AlibabaalihousenewhomecouponsubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_coupon_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomecouponsubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
