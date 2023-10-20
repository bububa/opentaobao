package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceCodeConsumeAPIResponse 天猫服务平台服务核销 API返回值
// tmall.service.code.consume
//
// 天猫服务平台－服务核销
type TmallServiceCodeConsumeAPIResponse struct {
	model.CommonResponse
	TmallServiceCodeConsumeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServiceCodeConsumeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServiceCodeConsumeAPIResponseModel).Reset()
}

// TmallServiceCodeConsumeAPIResponseModel is 天猫服务平台服务核销 成功返回结果
type TmallServiceCodeConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_service_code_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务工单ID
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServiceCodeConsumeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolTmallServiceCodeConsumeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServiceCodeConsumeAPIResponse)
	},
}

// GetTmallServiceCodeConsumeAPIResponse 从 sync.Pool 获取 TmallServiceCodeConsumeAPIResponse
func GetTmallServiceCodeConsumeAPIResponse() *TmallServiceCodeConsumeAPIResponse {
	return poolTmallServiceCodeConsumeAPIResponse.Get().(*TmallServiceCodeConsumeAPIResponse)
}

// ReleaseTmallServiceCodeConsumeAPIResponse 将 TmallServiceCodeConsumeAPIResponse 保存到 sync.Pool
func ReleaseTmallServiceCodeConsumeAPIResponse(v *TmallServiceCodeConsumeAPIResponse) {
	v.Reset()
	poolTmallServiceCodeConsumeAPIResponse.Put(v)
}
