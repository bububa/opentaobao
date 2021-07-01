package seaking

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSeakingAuthmachineapiAPIResponse
机翻Api授权 API返回值
alibaba.seaking.authmachineapi

机翻Api授权 */
type AlibabaSeakingAuthmachineapiAPIResponse struct {
	model.CommonResponse
	AlibabaSeakingAuthmachineapiAPIResponseModel
}

// AlibabaSeakingAuthmachineapiAPIResponseModel is 机翻Api授权 成功返回结果
type AlibabaSeakingAuthmachineapiAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_authmachineapi_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
