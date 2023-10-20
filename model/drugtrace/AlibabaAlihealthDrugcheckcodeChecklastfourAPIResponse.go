package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcheckcodechecklastfourAPIResponse 校验追溯码的后4位是否正确 API返回值
// alibaba.alihealth.drugcheckcode.checklastfour
//
// 校验追溯码的后4位是否正确
type AlibabaalihealthdrugcheckcodechecklastfourAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugcheckcodechecklastfourAPIResponseModel
}

// AlibabaalihealthdrugcheckcodechecklastfourAPIResponseModel is 校验追溯码的后4位是否正确 成功返回结果
type AlibabaalihealthdrugcheckcodechecklastfourAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcheckcode_checklastfour_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用成功
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 调用失败
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 对象
	SuccessI bool `json:"success_i,omitempty" xml:"success_i,omitempty"`
	// 是否正确
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
