package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushvenueAPIResponse 大麦换验平台-第三方对外开放-场馆接口pushVenue API返回值
// alibaba.damai.mev.open.pushvenue
//
// 开放接口推送场馆
type AlibabaDamaiMevOpenPushvenueAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenPushvenueAPIResponseModel
}

// AlibabaDamaiMevOpenPushvenueAPIResponseModel is 大麦换验平台-第三方对外开放-场馆接口pushVenue 成功返回结果
type AlibabaDamaiMevOpenPushvenueAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_pushvenue_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenPushvenueResult `json:"result,omitempty" xml:"result,omitempty"`
}
