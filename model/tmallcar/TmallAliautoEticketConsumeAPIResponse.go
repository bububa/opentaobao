package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoEticketConsumeAPIResponse 天猫汽车二轮车电子凭证核销 API返回值
// tmall.aliauto.eticket.consume
//
// 天猫汽车二轮车行业门店电子凭证核销
type TmallAliautoEticketConsumeAPIResponse struct {
	model.CommonResponse
	TmallAliautoEticketConsumeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAliautoEticketConsumeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAliautoEticketConsumeAPIResponseModel).Reset()
}

// TmallAliautoEticketConsumeAPIResponseModel is 天猫汽车二轮车电子凭证核销 成功返回结果
type TmallAliautoEticketConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_eticket_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的数据实体
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TmallAliautoEticketConsumeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = false
}

var poolTmallAliautoEticketConsumeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAliautoEticketConsumeAPIResponse)
	},
}

// GetTmallAliautoEticketConsumeAPIResponse 从 sync.Pool 获取 TmallAliautoEticketConsumeAPIResponse
func GetTmallAliautoEticketConsumeAPIResponse() *TmallAliautoEticketConsumeAPIResponse {
	return poolTmallAliautoEticketConsumeAPIResponse.Get().(*TmallAliautoEticketConsumeAPIResponse)
}

// ReleaseTmallAliautoEticketConsumeAPIResponse 将 TmallAliautoEticketConsumeAPIResponse 保存到 sync.Pool
func ReleaseTmallAliautoEticketConsumeAPIResponse(v *TmallAliautoEticketConsumeAPIResponse) {
	v.Reset()
	poolTmallAliautoEticketConsumeAPIResponse.Put(v)
}
