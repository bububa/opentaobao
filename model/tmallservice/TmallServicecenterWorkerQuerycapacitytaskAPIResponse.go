package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkerQuerycapacitytaskAPIResponse 查询需求容量 API返回值
// tmall.servicecenter.worker.querycapacitytask
//
// 查询需求容量
type TmallServicecenterWorkerQuerycapacitytaskAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkerQuerycapacitytaskAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkerQuerycapacitytaskAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkerQuerycapacitytaskAPIResponseModel).Reset()
}

// TmallServicecenterWorkerQuerycapacitytaskAPIResponseModel is 查询需求容量 成功返回结果
type TmallServicecenterWorkerQuerycapacitytaskAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_worker_querycapacitytask_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ResultBase
	ResultBase *ResultBase `json:"result_base,omitempty" xml:"result_base,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkerQuerycapacitytaskAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultBase = nil
}

var poolTmallServicecenterWorkerQuerycapacitytaskAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkerQuerycapacitytaskAPIResponse)
	},
}

// GetTmallServicecenterWorkerQuerycapacitytaskAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkerQuerycapacitytaskAPIResponse
func GetTmallServicecenterWorkerQuerycapacitytaskAPIResponse() *TmallServicecenterWorkerQuerycapacitytaskAPIResponse {
	return poolTmallServicecenterWorkerQuerycapacitytaskAPIResponse.Get().(*TmallServicecenterWorkerQuerycapacitytaskAPIResponse)
}

// ReleaseTmallServicecenterWorkerQuerycapacitytaskAPIResponse 将 TmallServicecenterWorkerQuerycapacitytaskAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkerQuerycapacitytaskAPIResponse(v *TmallServicecenterWorkerQuerycapacitytaskAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkerQuerycapacitytaskAPIResponse.Put(v)
}
