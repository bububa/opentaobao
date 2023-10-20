package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenpushvenueAPIResponse 大麦换验平台-第三方对外开放-场馆接口pushVenue API返回值
// alibaba.damai.mev.open.pushvenue
//
// 开放接口推送场馆
type AlibabadamaimevopenpushvenueAPIResponse struct {
	model.CommonResponse
	AlibabadamaimevopenpushvenueAPIResponseModel
}

// AlibabadamaimevopenpushvenueAPIResponseModel is 大麦换验平台-第三方对外开放-场馆接口pushVenue 成功返回结果
type AlibabadamaimevopenpushvenueAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_pushvenue_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabadamaimevopenpushvenueResult `json:"result,omitempty" xml:"result,omitempty"`
}
