package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenpushstandAPIResponse 大麦换验平台-第三方对外开放-看台接口pushStand API返回值
// alibaba.damai.mev.open.pushstand
//
// pushStand
type AlibabadamaimevopenpushstandAPIResponse struct {
	model.CommonResponse
	AlibabadamaimevopenpushstandAPIResponseModel
}

// AlibabadamaimevopenpushstandAPIResponseModel is 大麦换验平台-第三方对外开放-看台接口pushStand 成功返回结果
type AlibabadamaimevopenpushstandAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_pushstand_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabadamaimevopenpushstandResult `json:"result,omitempty" xml:"result,omitempty"`
}
