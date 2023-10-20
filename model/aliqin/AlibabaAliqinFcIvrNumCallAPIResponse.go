package aliqin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinfcivrnumcallAPIResponse ivr呼叫 API返回值
// alibaba.aliqin.fc.ivr.num.call
//
// ivr呼叫
type AlibabaaliqinfcivrnumcallAPIResponse struct {
	model.CommonResponse
	AlibabaaliqinfcivrnumcallAPIResponseModel
}

// AlibabaaliqinfcivrnumcallAPIResponseModel is ivr呼叫 成功返回结果
type AlibabaaliqinfcivrnumcallAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_ivr_num_call_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaaliqinfcivrnumcallResult `json:"result,omitempty" xml:"result,omitempty"`
}
