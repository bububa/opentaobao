package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardCompleteAPIResponse 工单完结 API返回值
// tmall.servicecenter.workcard.complete
//
// 工单完结
type TmallServicecenterWorkcardCompleteAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardCompleteAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardCompleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardCompleteAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardCompleteAPIResponseModel is 工单完结 成功返回结果
type TmallServicecenterWorkcardCompleteAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_complete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardCompleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardCompleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardCompleteAPIResponse)
	},
}

// GetTmallServicecenterWorkcardCompleteAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardCompleteAPIResponse
func GetTmallServicecenterWorkcardCompleteAPIResponse() *TmallServicecenterWorkcardCompleteAPIResponse {
	return poolTmallServicecenterWorkcardCompleteAPIResponse.Get().(*TmallServicecenterWorkcardCompleteAPIResponse)
}

// ReleaseTmallServicecenterWorkcardCompleteAPIResponse 将 TmallServicecenterWorkcardCompleteAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardCompleteAPIResponse(v *TmallServicecenterWorkcardCompleteAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardCompleteAPIResponse.Put(v)
}
