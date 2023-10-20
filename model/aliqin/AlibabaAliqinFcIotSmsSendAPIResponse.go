package aliqin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinfciotsmssendAPIResponse 物联网短信发送 API返回值
// alibaba.aliqin.fc.iot.sms.send
//
// 发送物联网短信，只允许使用物联网短信模板
type AlibabaaliqinfciotsmssendAPIResponse struct {
	model.CommonResponse
	AlibabaaliqinfciotsmssendAPIResponseModel
}

// AlibabaaliqinfciotsmssendAPIResponseModel is 物联网短信发送 成功返回结果
type AlibabaaliqinfciotsmssendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_sms_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaaliqinfciotsmssendBizResult `json:"result,omitempty" xml:"result,omitempty"`
}
