package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmemberqrcodeidentifyAPIResponse 扫码识别会员接口 API返回值
// alibaba.wdk.member.qrcode.identify
//
// 根据用户输入的付款码（支付宝、盒马、淘宝）、商家等信息，查询当前用户的基本信息及对应会员卡信息
type AlibabawdkmemberqrcodeidentifyAPIResponse struct {
	model.CommonResponse
	AlibabawdkmemberqrcodeidentifyAPIResponseModel
}

// AlibabawdkmemberqrcodeidentifyAPIResponseModel is 扫码识别会员接口 成功返回结果
type AlibabawdkmemberqrcodeidentifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_member_qrcode_identify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabawdkmemberqrcodeidentifyMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
