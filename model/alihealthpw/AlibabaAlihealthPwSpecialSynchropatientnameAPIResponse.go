package alihealthpw

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpwspecialsynchropatientnameAPIResponse 同步患者姓名至阿里健康 API返回值
// alibaba.alihealth.pw.special.synchropatientname
//
// 同步患者姓名至阿里健康
type AlibabaalihealthpwspecialsynchropatientnameAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthpwspecialsynchropatientnameAPIResponseModel
}

// AlibabaalihealthpwspecialsynchropatientnameAPIResponseModel is 同步患者姓名至阿里健康 成功返回结果
type AlibabaalihealthpwspecialsynchropatientnameAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pw_special_synchropatientname_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResponseMessage `json:"result,omitempty" xml:"result,omitempty"`
}
