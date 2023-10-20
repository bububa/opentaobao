package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytcodereplaceAPIResponse 单码替换 API返回值
// alibaba.alihealth.drug.kyt.codereplace
//
// 单码替换
type AlibabaalihealthdrugkytcodereplaceAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytcodereplaceAPIResponseModel
}

// AlibabaalihealthdrugkytcodereplaceAPIResponseModel is 单码替换 成功返回结果
type AlibabaalihealthdrugkytcodereplaceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_codereplace_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回列表
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 调用是否成功(true 成功 false 失败)
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
