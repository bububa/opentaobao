package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardSuspendAPIResponse 工单挂起 API返回值
// tmall.servicecenter.workcard.suspend
//
// 工单挂起
type TmallServicecenterWorkcardSuspendAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardSuspendAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardSuspendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardSuspendAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardSuspendAPIResponseModel is 工单挂起 成功返回结果
type TmallServicecenterWorkcardSuspendAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_suspend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *TmallServicecenterWorkcardSuspendResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardSuspendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardSuspendAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardSuspendAPIResponse)
	},
}

// GetTmallServicecenterWorkcardSuspendAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardSuspendAPIResponse
func GetTmallServicecenterWorkcardSuspendAPIResponse() *TmallServicecenterWorkcardSuspendAPIResponse {
	return poolTmallServicecenterWorkcardSuspendAPIResponse.Get().(*TmallServicecenterWorkcardSuspendAPIResponse)
}

// ReleaseTmallServicecenterWorkcardSuspendAPIResponse 将 TmallServicecenterWorkcardSuspendAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardSuspendAPIResponse(v *TmallServicecenterWorkcardSuspendAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardSuspendAPIResponse.Put(v)
}
