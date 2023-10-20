package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenresetticketAPIResponse 大麦换验平台-第三方对外开放-票单接口resetTicket API返回值
// alibaba.damai.mev.open.resetticket
//
// 开放接口重打票
type AlibabadamaimevopenresetticketAPIResponse struct {
	model.CommonResponse
	AlibabadamaimevopenresetticketAPIResponseModel
}

// AlibabadamaimevopenresetticketAPIResponseModel is 大麦换验平台-第三方对外开放-票单接口resetTicket 成功返回结果
type AlibabadamaimevopenresetticketAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_resetticket_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabadamaimevopenresetticketResult `json:"result,omitempty" xml:"result,omitempty"`
}
