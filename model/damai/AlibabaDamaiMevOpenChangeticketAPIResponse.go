package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenchangeticketAPIResponse 大麦换验平台-第三方对外开放-票单接口changeTicket API返回值
// alibaba.damai.mev.open.changeticket
//
// 开放接口 换票
type AlibabadamaimevopenchangeticketAPIResponse struct {
	model.CommonResponse
	AlibabadamaimevopenchangeticketAPIResponseModel
}

// AlibabadamaimevopenchangeticketAPIResponseModel is 大麦换验平台-第三方对外开放-票单接口changeTicket 成功返回结果
type AlibabadamaimevopenchangeticketAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_changeticket_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabadamaimevopenchangeticketResult `json:"result,omitempty" xml:"result,omitempty"`
}
