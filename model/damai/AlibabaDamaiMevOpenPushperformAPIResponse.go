package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenpushperformAPIResponse 大麦换验平台-第三方对外开放-场次接口pushPerform API返回值
// alibaba.damai.mev.open.pushperform
//
// pushPerform
type AlibabadamaimevopenpushperformAPIResponse struct {
	model.CommonResponse
	AlibabadamaimevopenpushperformAPIResponseModel
}

// AlibabadamaimevopenpushperformAPIResponseModel is 大麦换验平台-第三方对外开放-场次接口pushPerform 成功返回结果
type AlibabadamaimevopenpushperformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_pushperform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabadamaimevopenpushperformResult `json:"result,omitempty" xml:"result,omitempty"`
}
