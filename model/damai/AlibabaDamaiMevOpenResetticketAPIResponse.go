package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenResetticketAPIResponse 大麦换验平台-第三方对外开放-票单接口resetTicket API返回值
// alibaba.damai.mev.open.resetticket
//
// 开放接口重打票
type AlibabaDamaiMevOpenResetticketAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenResetticketAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenResetticketAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenResetticketAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenResetticketAPIResponseModel is 大麦换验平台-第三方对外开放-票单接口resetTicket 成功返回结果
type AlibabaDamaiMevOpenResetticketAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_resetticket_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenResetticketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenResetticketAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenResetticketAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenResetticketAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenResetticketAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenResetticketAPIResponse
func GetAlibabaDamaiMevOpenResetticketAPIResponse() *AlibabaDamaiMevOpenResetticketAPIResponse {
	return poolAlibabaDamaiMevOpenResetticketAPIResponse.Get().(*AlibabaDamaiMevOpenResetticketAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenResetticketAPIResponse 将 AlibabaDamaiMevOpenResetticketAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenResetticketAPIResponse(v *AlibabaDamaiMevOpenResetticketAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenResetticketAPIResponse.Put(v)
}
