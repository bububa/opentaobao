package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkcoupontemplateupdateAPIResponse 优惠券模版修改 API返回值
// alibaba.wdk.coupon.template.update
//
// 优惠券模版修改
type AlibabawdkcoupontemplateupdateAPIResponse struct {
	model.CommonResponse
	AlibabawdkcoupontemplateupdateAPIResponseModel
}

// AlibabawdkcoupontemplateupdateAPIResponseModel is 优惠券模版修改 成功返回结果
type AlibabawdkcoupontemplateupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_coupon_template_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabawdkcoupontemplateupdateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
