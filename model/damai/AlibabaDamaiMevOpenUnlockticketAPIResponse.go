package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenUnlockticketAPIResponse 大麦换验平台-第三方对外开放-票单接口unlockTicket API返回值
// alibaba.damai.mev.open.unlockticket
//
// 开放接口 解锁票单
type AlibabaDamaiMevOpenUnlockticketAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenUnlockticketAPIResponseModel
}

// AlibabaDamaiMevOpenUnlockticketAPIResponseModel is 大麦换验平台-第三方对外开放-票单接口unlockTicket 成功返回结果
type AlibabaDamaiMevOpenUnlockticketAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_unlockticket_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenUnlockticketResult `json:"result,omitempty" xml:"result,omitempty"`
}
