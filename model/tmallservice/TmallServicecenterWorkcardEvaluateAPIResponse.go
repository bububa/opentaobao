package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardEvaluateAPIResponse 服务商反馈鉴定结果 API返回值
// tmall.servicecenter.workcard.evaluate
//
// 服务商反馈鉴定结果
type TmallServicecenterWorkcardEvaluateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardEvaluateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardEvaluateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardEvaluateAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardEvaluateAPIResponseModel is 服务商反馈鉴定结果 成功返回结果
type TmallServicecenterWorkcardEvaluateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_evaluate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值Result
	Result *TmallServicecenterWorkcardEvaluateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardEvaluateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardEvaluateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardEvaluateAPIResponse)
	},
}

// GetTmallServicecenterWorkcardEvaluateAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardEvaluateAPIResponse
func GetTmallServicecenterWorkcardEvaluateAPIResponse() *TmallServicecenterWorkcardEvaluateAPIResponse {
	return poolTmallServicecenterWorkcardEvaluateAPIResponse.Get().(*TmallServicecenterWorkcardEvaluateAPIResponse)
}

// ReleaseTmallServicecenterWorkcardEvaluateAPIResponse 将 TmallServicecenterWorkcardEvaluateAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardEvaluateAPIResponse(v *TmallServicecenterWorkcardEvaluateAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardEvaluateAPIResponse.Put(v)
}
