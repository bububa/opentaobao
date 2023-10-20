package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenlockticketAPIResponse 大麦换验平台-第三方对外开放-票单接口lockTicket API返回值
// alibaba.damai.mev.open.lockticket
//
// 开放接口 冻结票单
type AlibabadamaimevopenlockticketAPIResponse struct {
	model.CommonResponse
	AlibabadamaimevopenlockticketAPIResponseModel
}

// AlibabadamaimevopenlockticketAPIResponseModel is 大麦换验平台-第三方对外开放-票单接口lockTicket 成功返回结果
type AlibabadamaimevopenlockticketAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_lockticket_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabadamaimevopenlockticketResult `json:"result,omitempty" xml:"result,omitempty"`
}
