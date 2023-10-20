package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenChangeticketAPIResponse 大麦换验平台-第三方对外开放-票单接口changeTicket API返回值
// alibaba.damai.mev.open.changeticket
//
// 开放接口 换票
type AlibabaDamaiMevOpenChangeticketAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenChangeticketAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenChangeticketAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenChangeticketAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenChangeticketAPIResponseModel is 大麦换验平台-第三方对外开放-票单接口changeTicket 成功返回结果
type AlibabaDamaiMevOpenChangeticketAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_changeticket_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenChangeticketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenChangeticketAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenChangeticketAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenChangeticketAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenChangeticketAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenChangeticketAPIResponse
func GetAlibabaDamaiMevOpenChangeticketAPIResponse() *AlibabaDamaiMevOpenChangeticketAPIResponse {
	return poolAlibabaDamaiMevOpenChangeticketAPIResponse.Get().(*AlibabaDamaiMevOpenChangeticketAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenChangeticketAPIResponse 将 AlibabaDamaiMevOpenChangeticketAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenChangeticketAPIResponse(v *AlibabaDamaiMevOpenChangeticketAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenChangeticketAPIResponse.Put(v)
}
