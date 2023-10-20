package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodekytwescheckcoderelationAPIResponse 检查输入的码之间是否有上下级关系 API返回值
// alibaba.alihealth.drug.code.kyt.wes.checkcoderelation
//
// 检查输入的码之间是否有上下级关系
type AlibabaalihealthdrugcodekytwescheckcoderelationAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugcodekytwescheckcoderelationAPIResponseModel
}

// AlibabaalihealthdrugcodekytwescheckcoderelationAPIResponseModel is 检查输入的码之间是否有上下级关系 成功返回结果
type AlibabaalihealthdrugcodekytwescheckcoderelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_wes_checkcoderelation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的码的结果
	Model []WesCodeRelationDto `json:"model,omitempty" xml:"model>wes_code_relation_dto,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 调用是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
