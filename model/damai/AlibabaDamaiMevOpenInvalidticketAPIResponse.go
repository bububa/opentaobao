package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenInvalidticketAPIResponse 大麦换验平台-第三方对外开放-票单接口invalidTicket API返回值
// alibaba.damai.mev.open.invalidticket
//
// 开放接口 使票无效
type AlibabaDamaiMevOpenInvalidticketAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenInvalidticketAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenInvalidticketAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenInvalidticketAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenInvalidticketAPIResponseModel is 大麦换验平台-第三方对外开放-票单接口invalidTicket 成功返回结果
type AlibabaDamaiMevOpenInvalidticketAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_invalidticket_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenInvalidticketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenInvalidticketAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenInvalidticketAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenInvalidticketAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenInvalidticketAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenInvalidticketAPIResponse
func GetAlibabaDamaiMevOpenInvalidticketAPIResponse() *AlibabaDamaiMevOpenInvalidticketAPIResponse {
	return poolAlibabaDamaiMevOpenInvalidticketAPIResponse.Get().(*AlibabaDamaiMevOpenInvalidticketAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenInvalidticketAPIResponse 将 AlibabaDamaiMevOpenInvalidticketAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenInvalidticketAPIResponse(v *AlibabaDamaiMevOpenInvalidticketAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenInvalidticketAPIResponse.Put(v)
}
