package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinfcvoicenumcancelcallAPIResponse 取消呼叫 API返回值
// alibaba.aliqin.fc.voice.num.cancelcall
//
// 当通话通过阿里大于呼出后可以通过调用这个接口取消本次通话
type AlibabaaliqinfcvoicenumcancelcallAPIResponse struct {
	model.CommonResponse
	AlibabaaliqinfcvoicenumcancelcallAPIResponseModel
}

// AlibabaaliqinfcvoicenumcancelcallAPIResponseModel is 取消呼叫 成功返回结果
type AlibabaaliqinfcvoicenumcancelcallAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_voice_num_cancelcall_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaaliqinfcvoicenumcancelcallBizResult `json:"result,omitempty" xml:"result,omitempty"`
}
