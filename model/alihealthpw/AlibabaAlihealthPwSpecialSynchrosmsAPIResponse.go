package alihealthpw

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwSpecialSynchrosmsAPIResponse 同步短信信息至阿里健康 API返回值
// alibaba.alihealth.pw.special.synchrosms
//
// 同步短信信息至阿里健康
type AlibabaAlihealthPwSpecialSynchrosmsAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthPwSpecialSynchrosmsAPIResponseModel
}

// AlibabaAlihealthPwSpecialSynchrosmsAPIResponseModel is 同步短信信息至阿里健康 成功返回结果
type AlibabaAlihealthPwSpecialSynchrosmsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pw_special_synchrosms_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResponseMessage `json:"result,omitempty" xml:"result,omitempty"`
}
