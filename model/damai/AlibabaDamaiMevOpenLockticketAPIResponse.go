package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenLockticketAPIResponse 大麦换验平台-第三方对外开放-票单接口lockTicket API返回值
// alibaba.damai.mev.open.lockticket
//
// 开放接口 冻结票单
type AlibabaDamaiMevOpenLockticketAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenLockticketAPIResponseModel
}

// AlibabaDamaiMevOpenLockticketAPIResponseModel is 大麦换验平台-第三方对外开放-票单接口lockTicket 成功返回结果
type AlibabaDamaiMevOpenLockticketAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_lockticket_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenLockticketResult `json:"result,omitempty" xml:"result,omitempty"`
}
