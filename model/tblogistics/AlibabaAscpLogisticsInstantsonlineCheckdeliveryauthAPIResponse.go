package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIResponse 同城配送在线下单检查授权 API返回值
// alibaba.ascp.logistics.instantsonline.checkdeliveryauth
//
// 同城配送在线下单检查授权
type AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIResponse struct {
	model.CommonResponse
	AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIResponseModel
}

// AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIResponseModel is 同城配送在线下单检查授权 成功返回结果
type AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_instantsonline_checkdeliveryauth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaascplogisticsinstantsonlinecheckdeliveryauthTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
