package middleclaims

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMiddleClaimsacceptReceiveAPIResponse 国际化中台服务域接收保险公司理赔受理结果 API返回值
// alibaba.middle.claimsaccept.receive
//
// 国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔受理结果的处理后，将该结果返回至服务域
type AlibabaMiddleClaimsacceptReceiveAPIResponse struct {
	model.CommonResponse
	AlibabaMiddleClaimsacceptReceiveAPIResponseModel
}

// AlibabaMiddleClaimsacceptReceiveAPIResponseModel is 国际化中台服务域接收保险公司理赔受理结果 成功返回结果
type AlibabaMiddleClaimsacceptReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_middle_claimsaccept_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaMiddleClaimsacceptReceiveResult `json:"result,omitempty" xml:"result,omitempty"`
}
