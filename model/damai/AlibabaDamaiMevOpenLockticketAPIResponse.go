package damai

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenLockticketAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenLockticketAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenLockticketAPIResponseModel is 大麦换验平台-第三方对外开放-票单接口lockTicket 成功返回结果
type AlibabaDamaiMevOpenLockticketAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_lockticket_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenLockticketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenLockticketAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenLockticketAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenLockticketAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenLockticketAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenLockticketAPIResponse
func GetAlibabaDamaiMevOpenLockticketAPIResponse() *AlibabaDamaiMevOpenLockticketAPIResponse {
	return poolAlibabaDamaiMevOpenLockticketAPIResponse.Get().(*AlibabaDamaiMevOpenLockticketAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenLockticketAPIResponse 将 AlibabaDamaiMevOpenLockticketAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenLockticketAPIResponse(v *AlibabaDamaiMevOpenLockticketAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenLockticketAPIResponse.Put(v)
}
