package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenWithdrawticketAPIResponse 大麦换验平台-第三方对外开放-票单接口withdrawTicket API返回值
// alibaba.damai.mev.open.withdrawticket
//
// 开放接口退票
type AlibabaDamaiMevOpenWithdrawticketAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenWithdrawticketAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenWithdrawticketAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenWithdrawticketAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenWithdrawticketAPIResponseModel is 大麦换验平台-第三方对外开放-票单接口withdrawTicket 成功返回结果
type AlibabaDamaiMevOpenWithdrawticketAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_withdrawticket_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenWithdrawticketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenWithdrawticketAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenWithdrawticketAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenWithdrawticketAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenWithdrawticketAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenWithdrawticketAPIResponse
func GetAlibabaDamaiMevOpenWithdrawticketAPIResponse() *AlibabaDamaiMevOpenWithdrawticketAPIResponse {
	return poolAlibabaDamaiMevOpenWithdrawticketAPIResponse.Get().(*AlibabaDamaiMevOpenWithdrawticketAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenWithdrawticketAPIResponse 将 AlibabaDamaiMevOpenWithdrawticketAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenWithdrawticketAPIResponse(v *AlibabaDamaiMevOpenWithdrawticketAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenWithdrawticketAPIResponse.Put(v)
}
