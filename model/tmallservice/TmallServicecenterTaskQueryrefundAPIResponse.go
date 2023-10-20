package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterTaskQueryrefundAPIResponse 查询任务类工单是否退款 API返回值
// tmall.servicecenter.task.queryrefund
//
// 查询任务类工单是否退款
type TmallServicecenterTaskQueryrefundAPIResponse struct {
	model.CommonResponse
	TmallServicecenterTaskQueryrefundAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterTaskQueryrefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterTaskQueryrefundAPIResponseModel).Reset()
}

// TmallServicecenterTaskQueryrefundAPIResponseModel is 查询任务类工单是否退款 成功返回结果
type TmallServicecenterTaskQueryrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_task_queryrefund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterTaskQueryrefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterTaskQueryrefundAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterTaskQueryrefundAPIResponse)
	},
}

// GetTmallServicecenterTaskQueryrefundAPIResponse 从 sync.Pool 获取 TmallServicecenterTaskQueryrefundAPIResponse
func GetTmallServicecenterTaskQueryrefundAPIResponse() *TmallServicecenterTaskQueryrefundAPIResponse {
	return poolTmallServicecenterTaskQueryrefundAPIResponse.Get().(*TmallServicecenterTaskQueryrefundAPIResponse)
}

// ReleaseTmallServicecenterTaskQueryrefundAPIResponse 将 TmallServicecenterTaskQueryrefundAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterTaskQueryrefundAPIResponse(v *TmallServicecenterTaskQueryrefundAPIResponse) {
	v.Reset()
	poolTmallServicecenterTaskQueryrefundAPIResponse.Put(v)
}
