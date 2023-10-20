package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkcouponspreadapplyAPIResponse 普通发券 API返回值
// alibaba.wdk.coupon.spread.apply
//
// 优惠券发放
type AlibabawdkcouponspreadapplyAPIResponse struct {
	model.CommonResponse
	AlibabawdkcouponspreadapplyAPIResponseModel
}

// AlibabawdkcouponspreadapplyAPIResponseModel is 普通发券 成功返回结果
type AlibabawdkcouponspreadapplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_coupon_spread_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabawdkcouponspreadapplyApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
