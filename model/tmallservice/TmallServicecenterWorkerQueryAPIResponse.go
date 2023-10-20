package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkerQueryAPIResponse 工人信息查询 API返回值
// tmall.servicecenter.worker.query
//
// 查询服务商对应的工人信息
type TmallServicecenterWorkerQueryAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkerQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkerQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkerQueryAPIResponseModel).Reset()
}

// TmallServicecenterWorkerQueryAPIResponseModel is 工人信息查询 成功返回结果
type TmallServicecenterWorkerQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_worker_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkerQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkerQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkerQueryAPIResponse)
	},
}

// GetTmallServicecenterWorkerQueryAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkerQueryAPIResponse
func GetTmallServicecenterWorkerQueryAPIResponse() *TmallServicecenterWorkerQueryAPIResponse {
	return poolTmallServicecenterWorkerQueryAPIResponse.Get().(*TmallServicecenterWorkerQueryAPIResponse)
}

// ReleaseTmallServicecenterWorkerQueryAPIResponse 将 TmallServicecenterWorkerQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkerQueryAPIResponse(v *TmallServicecenterWorkerQueryAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkerQueryAPIResponse.Put(v)
}
