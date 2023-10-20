package alihealthpw

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpwspecialsynchronodeAPIResponse 合作方同步状态至阿里健康 API返回值
// alibaba.alihealth.pw.special.synchronode
//
// 合作方同步状态至阿里健康
type AlibabaalihealthpwspecialsynchronodeAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthpwspecialsynchronodeAPIResponseModel
}

// AlibabaalihealthpwspecialsynchronodeAPIResponseModel is 合作方同步状态至阿里健康 成功返回结果
type AlibabaalihealthpwspecialsynchronodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pw_special_synchronode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResponseMessage `json:"result,omitempty" xml:"result,omitempty"`
}
