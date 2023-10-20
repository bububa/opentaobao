package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkcouponabandonAPIResponse 废券 API返回值
// alibaba.wdk.coupon.abandon
//
// 优惠券废弃
type AlibabawdkcouponabandonAPIResponse struct {
	model.CommonResponse
	AlibabawdkcouponabandonAPIResponseModel
}

// AlibabawdkcouponabandonAPIResponseModel is 废券 成功返回结果
type AlibabawdkcouponabandonAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_coupon_abandon_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabawdkcouponabandonApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
