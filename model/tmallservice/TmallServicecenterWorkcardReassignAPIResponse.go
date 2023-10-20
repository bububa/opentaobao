package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardReassignAPIResponse 工单改派门店 API返回值
// tmall.servicecenter.workcard.reassign
//
// 工单改派门店
type TmallServicecenterWorkcardReassignAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardReassignAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardReassignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardReassignAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardReassignAPIResponseModel is 工单改派门店 成功返回结果
type TmallServicecenterWorkcardReassignAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_reassign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *TmallServicecenterWorkcardReassignResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardReassignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardReassignAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardReassignAPIResponse)
	},
}

// GetTmallServicecenterWorkcardReassignAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardReassignAPIResponse
func GetTmallServicecenterWorkcardReassignAPIResponse() *TmallServicecenterWorkcardReassignAPIResponse {
	return poolTmallServicecenterWorkcardReassignAPIResponse.Get().(*TmallServicecenterWorkcardReassignAPIResponse)
}

// ReleaseTmallServicecenterWorkcardReassignAPIResponse 将 TmallServicecenterWorkcardReassignAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardReassignAPIResponse(v *TmallServicecenterWorkcardReassignAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardReassignAPIResponse.Put(v)
}
