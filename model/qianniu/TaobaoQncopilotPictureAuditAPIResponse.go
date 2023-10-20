package qianniu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQncopilotPictureAuditAPIResponse AIGC创作图片审核 API返回值
// taobao.qncopilot.picture.audit
//
// AIGC创作图片审核
type TaobaoQncopilotPictureAuditAPIResponse struct {
	model.CommonResponse
	TaobaoQncopilotPictureAuditAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQncopilotPictureAuditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQncopilotPictureAuditAPIResponseModel).Reset()
}

// TaobaoQncopilotPictureAuditAPIResponseModel is AIGC创作图片审核 成功返回结果
type TaobaoQncopilotPictureAuditAPIResponseModel struct {
	XMLName xml.Name `xml:"qncopilot_picture_audit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 审核结果
	QnCopilotResultDo *QnCopilotResultDo `json:"qn_copilot_result_do,omitempty" xml:"qn_copilot_result_do,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQncopilotPictureAuditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.QnCopilotResultDo = nil
}

var poolTaobaoQncopilotPictureAuditAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQncopilotPictureAuditAPIResponse)
	},
}

// GetTaobaoQncopilotPictureAuditAPIResponse 从 sync.Pool 获取 TaobaoQncopilotPictureAuditAPIResponse
func GetTaobaoQncopilotPictureAuditAPIResponse() *TaobaoQncopilotPictureAuditAPIResponse {
	return poolTaobaoQncopilotPictureAuditAPIResponse.Get().(*TaobaoQncopilotPictureAuditAPIResponse)
}

// ReleaseTaobaoQncopilotPictureAuditAPIResponse 将 TaobaoQncopilotPictureAuditAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQncopilotPictureAuditAPIResponse(v *TaobaoQncopilotPictureAuditAPIResponse) {
	v.Reset()
	poolTaobaoQncopilotPictureAuditAPIResponse.Put(v)
}
