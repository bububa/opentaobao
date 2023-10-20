package damai

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenUnlockticketAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenUnlockticketAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenUnlockticketAPIResponseModel is 大麦换验平台-第三方对外开放-票单接口unlockTicket 成功返回结果
type AlibabaDamaiMevOpenUnlockticketAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_unlockticket_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenUnlockticketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenUnlockticketAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenUnlockticketAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenUnlockticketAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenUnlockticketAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenUnlockticketAPIResponse
func GetAlibabaDamaiMevOpenUnlockticketAPIResponse() *AlibabaDamaiMevOpenUnlockticketAPIResponse {
	return poolAlibabaDamaiMevOpenUnlockticketAPIResponse.Get().(*AlibabaDamaiMevOpenUnlockticketAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenUnlockticketAPIResponse 将 AlibabaDamaiMevOpenUnlockticketAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenUnlockticketAPIResponse(v *AlibabaDamaiMevOpenUnlockticketAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenUnlockticketAPIResponse.Put(v)
}
