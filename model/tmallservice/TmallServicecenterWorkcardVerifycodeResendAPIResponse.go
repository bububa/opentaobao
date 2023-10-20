package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardVerifycodeResendAPIResponse 重发核销码 API返回值
// tmall.servicecenter.workcard.verifycode.resend
//
// 重发核销码
type TmallServicecenterWorkcardVerifycodeResendAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardVerifycodeResendAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardVerifycodeResendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardVerifycodeResendAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardVerifycodeResendAPIResponseModel is 重发核销码 成功返回结果
type TmallServicecenterWorkcardVerifycodeResendAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_verifycode_resend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *TmallServicecenterWorkcardVerifycodeResendResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardVerifycodeResendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardVerifycodeResendAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardVerifycodeResendAPIResponse)
	},
}

// GetTmallServicecenterWorkcardVerifycodeResendAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardVerifycodeResendAPIResponse
func GetTmallServicecenterWorkcardVerifycodeResendAPIResponse() *TmallServicecenterWorkcardVerifycodeResendAPIResponse {
	return poolTmallServicecenterWorkcardVerifycodeResendAPIResponse.Get().(*TmallServicecenterWorkcardVerifycodeResendAPIResponse)
}

// ReleaseTmallServicecenterWorkcardVerifycodeResendAPIResponse 将 TmallServicecenterWorkcardVerifycodeResendAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardVerifycodeResendAPIResponse(v *TmallServicecenterWorkcardVerifycodeResendAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardVerifycodeResendAPIResponse.Put(v)
}
