package alihealthpw

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpwspecialsynchrosmsAPIResponse 同步短信信息至阿里健康 API返回值
// alibaba.alihealth.pw.special.synchrosms
//
// 同步短信信息至阿里健康
type AlibabaalihealthpwspecialsynchrosmsAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthpwspecialsynchrosmsAPIResponseModel
}

// AlibabaalihealthpwspecialsynchrosmsAPIResponseModel is 同步短信信息至阿里健康 成功返回结果
type AlibabaalihealthpwspecialsynchrosmsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pw_special_synchrosms_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResponseMessage `json:"result,omitempty" xml:"result,omitempty"`
}
