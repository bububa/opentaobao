package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenBatchpushticketAPIResponse 大麦换验平台-第三方对外开放-票单接口batchPushTicket API返回值
// alibaba.damai.mev.open.batchpushticket
//
// 批量推送票单
type AlibabaDamaiMevOpenBatchpushticketAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenBatchpushticketAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenBatchpushticketAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenBatchpushticketAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenBatchpushticketAPIResponseModel is 大麦换验平台-第三方对外开放-票单接口batchPushTicket 成功返回结果
type AlibabaDamaiMevOpenBatchpushticketAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_batchpushticket_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenBatchpushticketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenBatchpushticketAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenBatchpushticketAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenBatchpushticketAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenBatchpushticketAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenBatchpushticketAPIResponse
func GetAlibabaDamaiMevOpenBatchpushticketAPIResponse() *AlibabaDamaiMevOpenBatchpushticketAPIResponse {
	return poolAlibabaDamaiMevOpenBatchpushticketAPIResponse.Get().(*AlibabaDamaiMevOpenBatchpushticketAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenBatchpushticketAPIResponse 将 AlibabaDamaiMevOpenBatchpushticketAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenBatchpushticketAPIResponse(v *AlibabaDamaiMevOpenBatchpushticketAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenBatchpushticketAPIResponse.Put(v)
}
