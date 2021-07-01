package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMemberQrcodeIdentifyAPIResponse
扫码识别会员接口 API返回值
alibaba.wdk.member.qrcode.identify

根据用户输入的付款码（支付宝、盒马、淘宝）、商家等信息，查询当前用户的基本信息及对应会员卡信息 */
type AlibabaWdkMemberQrcodeIdentifyAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMemberQrcodeIdentifyAPIResponseModel
}

// AlibabaWdkMemberQrcodeIdentifyAPIResponseModel is 扫码识别会员接口 成功返回结果
type AlibabaWdkMemberQrcodeIdentifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_member_qrcode_identify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkMemberQrcodeIdentifyMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
