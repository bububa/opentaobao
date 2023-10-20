package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardAssignworkerAPIResponse 服务商分派工人 API返回值
// tmall.servicecenter.workcard.assignworker
//
// 服务商调用该接口分派工人给具体的工单
type TmallServicecenterWorkcardAssignworkerAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardAssignworkerAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardAssignworkerAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardAssignworkerAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardAssignworkerAPIResponseModel is 服务商分派工人 成功返回结果
type TmallServicecenterWorkcardAssignworkerAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_assignworker_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// -
	Result *TmallServicecenterWorkcardAssignworkerResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardAssignworkerAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardAssignworkerAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardAssignworkerAPIResponse)
	},
}

// GetTmallServicecenterWorkcardAssignworkerAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardAssignworkerAPIResponse
func GetTmallServicecenterWorkcardAssignworkerAPIResponse() *TmallServicecenterWorkcardAssignworkerAPIResponse {
	return poolTmallServicecenterWorkcardAssignworkerAPIResponse.Get().(*TmallServicecenterWorkcardAssignworkerAPIResponse)
}

// ReleaseTmallServicecenterWorkcardAssignworkerAPIResponse 将 TmallServicecenterWorkcardAssignworkerAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardAssignworkerAPIResponse(v *TmallServicecenterWorkcardAssignworkerAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardAssignworkerAPIResponse.Put(v)
}
